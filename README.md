# testpool

test grpc pool and streams pool

#### 1

pool: size = 200    

ttl = 1*time.Minute       

concurrency =1000

grpc pool result:

```go
tcp connections: 0 
tcp connections: 0 
tcp connections: 5 
tcp connections: 8 
tcp connections: 8 
tcp connections: 38 
tcp connections: 44 
tcp connections: 87 
tcp connections: 156 
tcp connections: 156 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 469 ms
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 10    

concurrency =1000

streams pool result:

```go
tcp connections: 91 
tcp connections: 102 
tcp connections: 131 
tcp connections: 134 
tcp connections: 141 
tcp connections: 141 
finish, use time: 150 ms 
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 1   

concurrency =1000

streams pool result:

```go
tcp connections: 167 
tcp connections: 184 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 465 ms
```



#### 2

pool: size = 200    

ttl = 1*time.Minute      

concurrency =2000

grpc pool result:

```go
tcp connections: 0 
tcp connections: 0 
tcp connections: 5 
tcp connections: 5 
tcp connections: 5 
tcp connections: 5 
tcp connections: 10 
tcp connections: 11 
tcp connections: 31 
tcp connections: 52 
tcp connections: 104 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 552 ms
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 10    

concurrency =2000

streams pool result:

```go
tcp connections: 151 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 202 ms
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 1   

concurrency =2000

streams pool result:

```go
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 415 ms
```



#### 3

pool: size = 200    

ttl = 1*time.Minute      

concurrency =4000

grpc pool result:

```go
tcp connections: 0 
tcp connections: 0 
tcp connections: 1 
tcp connections: 1 
tcp connections: 1 
tcp connections: 1 
tcp connections: 1 
tcp connections: 1 
tcp connections: 1 
tcp connections: 1 
tcp connections: 0 
tcp connections: 1 
tcp connections: 24 
tcp connections: 67 
tcp connections: 75 
tcp connections: 104 
tcp connections: 104 
tcp connections: 193 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 1294 ms
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 10   

concurrency =4000

streams pool result:

```go
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 466 ms
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 1  

concurrency =4000

streams pool result:

```go
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 1210 ms
```



#### more about streams pool

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 20  

concurrency =4000

streams pool result:

```go
tcp connections: 179 
tcp connections: 179 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 479 ms
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 40 

concurrency =4000

streams pool result:

```go
tcp connections: 111 
tcp connections: 120 
tcp connections: 136 
tcp connections: 136 
tcp connections: 139 
tcp connections: 139 
tcp connections: 139 
tcp connections: 139 
finish, use time: 140 ms
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 80 

concurrency =4000

streams pool result:

```go
tcp connections: 56 
tcp connections: 69 
tcp connections: 79 
tcp connections: 79 
tcp connections: 79 
tcp connections: 99 
tcp connections: 99 
tcp connections: 99 
tcp connections: 99 
tcp connections: 99 
finish, use time: 140 ms
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 400 

concurrency =4000

streams pool result:

```go
tcp connections: 19 
tcp connections: 19 
tcp connections: 19 
tcp connections: 19 
tcp connections: 19 
finish, use time: 73 ms
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 10000  

concurrency =4000

streams pool result:

```go
tcp connections: 13 
tcp connections: 13 
tcp connections: 13 
tcp connections: 13 
finish, use time: 74 ms
```



#### concurrency = 10000

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 50   

concurrency =10000

streams pool result:

```go
tcp connections: 185 
tcp connections: 187 
tcp connections: 187 
tcp connections: 187 
tcp connections: 187 
tcp connections: 187 
tcp connections: 187 
tcp connections: 187 
tcp connections: 187 
tcp connections: 187 
finish, use time: 259 ms
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 1000 

concurrency =10000

streams pool result:

```go
tcp connections: 15 
tcp connections: 15 
tcp connections: 22 
tcp connections: 22 
tcp connections: 22 
tcp connections: 22 
tcp connections: 22 
tcp connections: 22 
tcp connections: 22 
tcp connections: 22 
tcp connections: 22 
finish, use time: 200 ms
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 10000  

concurrency =10000

streams pool result:

```go
tcp connections: 13 
tcp connections: 13 
tcp connections: 13 
tcp connections: 13 
tcp connections: 13 
tcp connections: 13 
tcp connections: 13 
tcp connections: 13 
tcp connections: 13 
tcp connections: 13 
tcp connections: 13 
tcp connections: 13 
tcp connections: 13 
tcp connections: 13 
finish, use time: 204 ms
```

