## cloudgo-data

>  这次作业主要是对比XORM和database/sql的速度，就结果而言，XORM要比database稍微慢点

###  Test

> 8080端口是使用database/sql服务端的端口，8081是使用XORM服务端的宽口
>
> 首先看xorm的

```shell
ab -n 1000 -c 100 -p ./postFIle  http://127.0.0.1:8081/service/userinfo
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8081

Document Path:          /service/userinfo
Document Length:        3939 bytes

Concurrency Level:      100
Time taken for tests:   2.154 seconds
Complete requests:      1000
Failed requests:        916
   (Connect: 0, Receive: 0, Length: 916, Exceptions: 0)
Non-2xx responses:      1000
Total transferred:      4055935 bytes
Total body sent:        184000
HTML transferred:       3939935 bytes
Requests per second:    880.51 [#/sec] (mean)
Time per request:       113.570 [ms] (mean)
Time per request:       1.136 [ms] (mean, across all concurrent requests)
Transfer rate:          3487.60 [Kbytes/sec] received
                        158.22 kb/s sent
                        3645.82 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   4.3      0      43
Processing:     2  109  31.3    105     239
Waiting:        2  109  31.2    105     228
Total:          2  110  32.7    106     239

Percentage of the requests served within a certain time (ms)
  50%    106
  66%    113
  75%    117
  80%    123
  90%    153
  95%    176
  98%    211
  99%    216
 100%    239 (longest request)
 
 
ab -n 1000 -c 100   http://127.0.0.1:8081/service/userinfo?userid=1
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8081

Document Path:          /service/userinfo?userid=1
Document Length:        102 bytes

Concurrency Level:      100
Time taken for tests:   3.127 seconds
Complete requests:      1000
Failed requests:        0
Non-2xx responses:      1000
Total transferred:      235000 bytes
HTML transferred:       102000 bytes
Requests per second:    462.05 [#/sec] (mean)
Time per request:       216.425 [ms] (mean)
Time per request:       2.164 [ms] (mean, across all concurrent requests)
Transfer rate:          106.04 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    4   8.1      0      40
Processing:     2  210 122.5    184     605
Waiting:        2  208 123.0    182     586
Total:          2  214 125.5    186     613

Percentage of the requests served within a certain time (ms)
  50%    186
  66%    215
  75%    258
  80%    272
  90%    386
  95%    536
  98%    577
  99%    581
 100%    613 (longest request)



```

> 接着看database/sql的

```shell
ab -n 1000 -c 100 -p ./postFIle  http://127.0.0.1:8080/service/userinfo
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /service/userinfo
Document Length:        3994 bytes

Concurrency Level:      100
Time taken for tests:   1.316 seconds
Complete requests:      1000
Failed requests:        998
   (Connect: 0, Receive: 0, Length: 998, Exceptions: 0)
Non-2xx responses:      1000
Total transferred:      4111919 bytes
Total body sent:        184000
HTML transferred:       3995919 bytes
Requests per second:    759.89 [#/sec] (mean)
Time per request:       131.597 [ms] (mean)
Time per request:       1.316 [ms] (mean, across all concurrent requests)
Transfer rate:          3051.39 [Kbytes/sec] received
                        136.54 kb/s sent
                        3187.94 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   5.4      0      26
Processing:     1  124  33.6    129     267
Waiting:        0  123  33.7    128     267
Total:          1  126  31.5    130     267

Percentage of the requests served within a certain time (ms)
  50%    130
  66%    140
  75%    148
  80%    152
  90%    160
  95%    165
  98%    171
  99%    176
 100%    267 (longest request)


ab -n 1000 -c 100   http://127.0.0.1:8080/service/userinfo?userid=1
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /service/userinfo?userid=1
Document Length:        4505 bytes

Concurrency Level:      100
Time taken for tests:   2.257 seconds
Complete requests:      1000
Failed requests:        0
Non-2xx responses:      1000
Total transferred:      4621000 bytes
HTML transferred:       4505000 bytes
Requests per second:    442.99 [#/sec] (mean)
Time per request:       225.738 [ms] (mean)
Time per request:       2.257 [ms] (mean, across all concurrent requests)
Transfer rate:          1999.09 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    5  16.9      0      84
Processing:     4  216  58.9    211     533
Waiting:        4  216  58.9    210     533
Total:          4  221  58.7    218     533

Percentage of the requests served within a certain time (ms)
  50%    218
  66%    235
  75%    246
  80%    255
  90%    303
  95%    346
  98%    364
  99%    367
 100%    533 (longest request)

```

