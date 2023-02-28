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

For some modules, you can write benchmark in test files. Such as redis's benchmark.

#### Pressure Test

You can test the QPS by some tools, such as `ab`.
