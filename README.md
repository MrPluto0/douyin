# Douyin Project

## Introduction

This is a backend project about simple-douyin of Bytedance, using `gin` `gorm` and so on.

## APIS

- [x] POST `/douyin/user/login`
- [x] POST `douyin/user/register`
- [x] GET `/douyin/user/`
- [x] GET `/douyin/feed`

## How to develop

### Core

The core code is in `/app`, which has four parts:

- `controller`: it serves as apis' handlers.
- `service`: it's main business logic.
- `models`: it's the database models and functions(`dao`) to operate data.
- `define`: it's the data structure for each request and response.

From top to bottom layers: `Router -> Controller -> Service(define) -> Dao -> Model`

### Database

Database is initialized in `/init/database.go`.

#### Mysql

The main database is `mysql`, storing all persistent data. Its control is integrated in `gorm`.

Tips: you can create your datatables in database by `gorm.AutoMigrate(...)`

#### Redis

The cache database is `redis` in which data react quickly. But it's just used in api `GET /douyin/user`.

You can run benchmark in `/tests/redis_test.go` to see the performance.

### Config

You can edit the `config.yaml` to add another setting, and use in project like this:

```go
// learn how to use viper in go
viper.Get("new_config")
```

### Test

#### Feature Test

For each finished apis, it needs to be tested features in `tests` folder.

For each utils, it needs to be tested basicly in the same level folder.

The standards are as follows:

- Test all error code's examples which the apis may response.
- Write notes above this line or subtest including name to explain this test.

#### Benchmark

For some modules, you can write benchmark in test files. Such as redis and mysql benchmark.

```bash
goos: linux
goarch: amd64
pkg: douyin/tests
cpu: Intel(R) Xeon(R) Platinum 8255C CPU @ 2.50GHz
BenchmarkRedis
BenchmarkRedis-4   	   10000	    112403 ns/op	     188 B/op	       7 allocs/op
BenchmarkMysql
BenchmarkMysql-4   	    3570	    300697 ns/op	    9013 B/op	     134 allocs/op
PASS
ok  	douyin/tests	2.291s
```

#### Pressure Test

You can test the QPS by some tools, such as `ab`/`wrk`/...

The folder `wrk` is several scripts of `lua`, used to presssure tests.

Run the command in bash: `wrk -t12 -c400 -d10s -T30 -s ./wrk/login.lua http://127.0.0.1:8080`

Here is the result:

```bash
Running 10s test @ http://127.0.0.1:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    45.82ms   19.46ms 229.79ms   80.51%
    Req/Sec   733.62    134.87     1.19k    81.42%
  87715 requests in 10.09s, 14.89MB read
Requests/sec:   8694.54
Transfer/sec:      1.48MB
```