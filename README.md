# testpool

test grpc pool and streams pool

totalConnection means total dial connection

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
finish, use time: 486 ms, totalConnection: 999
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 10    

concurrency =1000

streams pool result:

```go
tcp connections: 91 
tcp connections: 117 
tcp connections: 129 
tcp connections: 129 
tcp connections: 154 
tcp connections: 154 
finish, use time: 135 ms, totalConnection: 154
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 1   

concurrency =1000

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
finish, use time: 478 ms, totalConnection: 1000
```



#### 2

pool: size = 200    

ttl = 1*time.Minute      

concurrency =2000

grpc pool result:

```go
tcp connections: 0 
tcp connections: 0 
tcp connections: 0 
tcp connections: 0 
tcp connections: 0 
tcp connections: 0 
tcp connections: 0 
tcp connections: 0 
tcp connections: 57 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 577 ms, totalConnection: 2000
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 10    

concurrency =2000

streams pool result:

```go
tcp connections: 187 
tcp connections: 189 
tcp connections: 189 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 162 ms, totalConnection: 316
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
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 550 ms, totalConnection: 1997
```



#### 3

pool: size = 200    

ttl = 1*time.Minute      

concurrency =4000

grpc pool result:

```go
tcp connections: 0 
tcp connections: 0 
tcp connections: 0 
tcp connections: 1 
tcp connections: 1 
tcp connections: 1 
tcp connections: 1 
tcp connections: 0 
tcp connections: 0 
tcp connections: 27 
tcp connections: 67 
tcp connections: 68 
tcp connections: 75 
tcp connections: 76 
tcp connections: 78 
tcp connections: 108 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 1271 ms, totalConnection: 3994
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
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 589 ms, totalConnection: 2200
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
finish, use time: 1305 ms, totalConnection: 3999
```



#### more about streams pool

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 20  

concurrency =4000

streams pool result:

```go
tcp connections: 196 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
tcp connections: 200 
finish, use time: 202 ms, totalConnection: 307
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 40 

concurrency =4000

streams pool result:

```go
tcp connections: 100 
tcp connections: 116 
tcp connections: 129 
tcp connections: 145 
tcp connections: 173 
tcp connections: 174 
tcp connections: 174 
tcp connections: 174 
tcp connections: 174 
tcp connections: 174 
tcp connections: 174 
finish, use time: 142 ms, totalConnection: 174
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 80 

concurrency =4000

streams pool result:

```go
tcp connections: 66 
tcp connections: 80 
tcp connections: 84 
tcp connections: 84 
tcp connections: 84 
tcp connections: 84 
tcp connections: 84 
tcp connections: 84 
tcp connections: 84 
finish, use time: 133 ms, totalConnection: 84
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 400 

concurrency =4000

streams pool result:

```go
tcp connections: 10 
tcp connections: 21 
tcp connections: 26 
tcp connections: 26 
tcp connections: 26 
tcp connections: 26 
tcp connections: 26 
tcp connections: 26 
tcp connections: 26 
finish, use time: 97 ms, totalConnection: 26
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 10000  

concurrency =4000

streams pool result:

```go
tcp connections: 5 
tcp connections: 5 
tcp connections: 5 
tcp connections: 5 
tcp connections: 5 
tcp connections: 5 
tcp connections: 5 
tcp connections: 5 
tcp connections: 5 
tcp connections: 5 
finish, use time: 99 ms, totalConnection: 5
```



#### concurrency = 10000

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 50   

concurrency =10000

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
finish, use time: 293 ms, totalConnection: 291
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 1000 

concurrency =10000

streams pool result:

```go
tcp connections: 43 
tcp connections: 45 
tcp connections: 66 
tcp connections: 85 
tcp connections: 85 
tcp connections: 85 
tcp connections: 85 
tcp connections: 85 
tcp connections: 85 
tcp connections: 85 
tcp connections: 85 
tcp connections: 85 
finish, use time: 266 ms, totalConnection: 85
```

pool: size = 200    

ttl = 1*time.Minute   

maxStreams = 10000  

concurrency =10000

streams pool result:

```go
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
tcp connections: 9 
finish, use time: 249 ms, totalConnection: 9
```

