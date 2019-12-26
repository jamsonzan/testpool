package main

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pgrpc "google.golang.org/grpc"
	pb "learnGrpc/testpool/helloworld"
)

type Greeter struct {}

func (g *Greeter)SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error)  {
	return &pb.HelloReply{Message: "Hello!" + request.Name}, nil
}

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
	p := newPool(200, 1*time.Minute, 20, 10000)
	for i := 0; i < 10000; i++ {
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
			//fmt.Printf("r")
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
			fmt.Printf("tcp connections: %d \n", p.conns[l.Addr().String()].count)
			p.Unlock()
			end := time.Now().UnixNano() / 1e6
			fmt.Printf("finish, use time: %d ms\n", end-begin)
			return
		case <-ticker.C:
			p.Lock()
			fmt.Printf("tcp connections: %d \n", p.conns[l.Addr().String()].count)
			p.Unlock()
		}
	}
}


type pool struct {
	size int
	ttl  int64
	//  max streams on a *poolConn
	maxStreams int
	//  max idle conns
	maxIdle    int

	sync.Mutex
	conns map[string]*streamsPool
}

type streamsPool struct {
	//  head of list
	head   *poolConn
	//  the siza of list
	count  int
	//  idle conn
	idle   int
}

type poolConn struct {
	//  grpc conn
	*grpc.ClientConn
	err     error
	addr    string

	//  pool and streams pool
	pool    *pool
	sp      *streamsPool
	streams int
	created int64

	//  list
	pre     *poolConn
	next    *poolConn
	in      bool
}

func newPool(size int, ttl time.Duration, ms int, idle int) *pool {
	if ms <= 0 {
		ms = 1
	}
	if idle <= 0 {
		idle = 0
	}
	return &pool{
		size:  size,
		ttl:   int64(ttl.Seconds()),
		maxStreams: ms,
		maxIdle: idle,
		conns: make(map[string]*streamsPool),
	}
}

func (p *pool) getConn(addr string, opts ...grpc.DialOption) (*poolConn, error) {
	now := time.Now().Unix()
	p.Lock()
	sp, ok := p.conns[addr]
	if !ok {
		sp = &streamsPool{head:&poolConn{}, count:0, idle:0}
		p.conns[addr] = sp
	}
	// while we have conns check streams and then return one
	// otherwise we'll create a new conn
	conn := sp.head.next
	for conn != nil {
		// too many streams or too old
		if conn.streams >= p.maxStreams || now-conn.created > p.ttl {
			conn = conn.next
			continue
		}
		// a idle conn
		if conn.streams == 0 {
			sp.idle--
		}
		// we got a good conn, lets unlock and return it
		conn.streams++
		p.Unlock()
		return conn, nil
	}
	p.Unlock()

	// create new conn
	cc, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	conn = &poolConn{cc,nil,addr,p,sp,1,time.Now().Unix(), nil, nil, false}

	// add conn to streams pool
	p.Lock()
	if sp.count < p.size {
		conn.in = true
		addConnAfter(conn, sp.head)
	}
	p.Unlock()

	return conn, nil
}

func (p *pool) release(addr string, conn *poolConn, err error) {
	p.Lock()
	conn.streams--
	if conn.streams > 0 {
		p.Unlock()
		return
	}
	now := time.Now().Unix()
	p, sp, created := conn.pool, conn.sp, conn.created
	//  add conn to streams pool
	if !conn.in && sp.count < p.size {
		conn.in = true
		addConnAfter(conn, sp.head)
	}
	// it has errored or
	// too many idle conn or
	// too many conn
	// conn is too old
	if !conn.in || err != nil || sp.idle > p.maxIdle || sp.count > p.size || now-created > p.ttl {
		removeConn(conn)
		p.Unlock()
		conn.ClientConn.Close()
		return
	}
	sp.idle++
	p.Unlock()
	return
}

func (conn *poolConn)Close()  {
	conn.pool.release(conn.addr, conn, conn.err)
}

func removeConn(conn *poolConn)  {
	if conn.next == nil || conn.pre == nil{
		conn.pre = nil
		return
	}
	conn.pre.next = conn.next
	conn.next.pre = conn.pre
	conn.sp.count--
	return
}

func addConnAfter(conn *poolConn, after *poolConn)  {
	conn.next = after.next
	conn.pre = after
	if after.next != nil {
		after.next.pre = conn
	}
	after.next = conn
	conn.sp.count++
	return
}