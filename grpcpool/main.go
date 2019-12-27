package main

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pgrpc "google.golang.org/grpc"
	pb "testpool/helloworld"
)
type Greeter struct {}

func (g *Greeter)SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error)  {
	return &pb.HelloReply{Message: "Hello!" + request.Name}, nil
}

var totalConnection int

func main() {
	// setup server
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer l.Close()

	s := pgrpc.NewServer()
	pb.RegisterGreeterServer(s, new(Greeter))

	go s.Serve(l)
	defer s.Stop()

	wg := &sync.WaitGroup{}
	// zero pool
	begin := time.Now().UnixNano() / 1e6
	p := newPool(200, 1*time.Minute)
	for i := 0; i < 400; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// get a conn
			cc, err := p.getConn(l.Addr().String(), grpc.WithInsecure())
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			rsp := pb.HelloReply{}

			err = cc.Invoke(context.TODO(), "/helloworld.Greeter/SayHello", &pb.HelloRequest{Name: "John"}, &rsp)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if rsp.Message != "Hello!John" {
				fmt.Println("reply error")
				return
			}

			// release the conn
			p.release(l.Addr().String(), cc, nil)
		}()
	}

	done := make(chan int)
	go func() {
		wg.Wait()
		done <- 1
	}()

	ticker := time.NewTicker(10*time.Millisecond)
	defer ticker.Stop()
	for  {
		select {
		case <-done:
			p.Lock()
			fmt.Printf("tcp connections: %d \n", len(p.conns[l.Addr().String()]))
			p.Unlock()
			end := time.Now().UnixNano() / 1e6
			fmt.Printf("finish, use time: %d ms, totalConnection: %d\n", end-begin, totalConnection)
			return
		case <-ticker.C:
			p.Lock()
			fmt.Printf("tcp connections: %d \n", len(p.conns[l.Addr().String()]))
			p.Unlock()
		}
	}

}

//  pool
type pool struct {
	size int
	ttl  int64

	sync.Mutex
	conns map[string][]*poolConn
}

type poolConn struct {
	*grpc.ClientConn
	created int64
}

func newPool(size int, ttl time.Duration) *pool {
	return &pool{
		size:  size,
		ttl:   int64(ttl.Seconds()),
		conns: make(map[string][]*poolConn),
	}
}

func (p *pool) getConn(addr string, opts ...grpc.DialOption) (*poolConn, error) {
	p.Lock()
	conns := p.conns[addr]
	now := time.Now().Unix()

	// while we have conns check age and then return one
	// otherwise we'll create a new conn
	for len(conns) > 0 {
		conn := conns[len(conns)-1]
		conns = conns[:len(conns)-1]
		p.conns[addr] = conns

		// if conn is old kill it and move on
		if d := now - conn.created; d > p.ttl {
			conn.ClientConn.Close()
			continue
		}

		// we got a good conn, lets unlock and return it
		p.Unlock()

		return conn, nil
	}
	totalConnection++
	p.Unlock()

	// create new conn
	cc, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	return &poolConn{cc, time.Now().Unix()}, nil
}

func (p *pool) release(addr string, conn *poolConn, err error) {
	// don't store the conn if it has errored
	if err != nil {
		conn.ClientConn.Close()
		return
	}

	// otherwise put it back for reuse
	p.Lock()
	conns := p.conns[addr]
	if len(conns) >= p.size {
		p.Unlock()
		conn.ClientConn.Close()
		return
	}
	p.conns[addr] = append(conns, conn)
	p.Unlock()
}
