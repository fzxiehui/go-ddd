# DDD 项目构建

## 登录测试

```shell
curl -X POST -H "Content-Type: application/json" -d '{"username": "root", "password": "root"}' http://127.0.0.1:8080/auth/login
```

## 注册测试

```shell
curl -X POST -H "Content-Type: application/json" -d '{"username": "root", "password": "root"}' http://127.0.0.1:8080/auth/register
```


## wire

```shell
go install github.com/google/wire/cmd/wire@v0.6.0
go get github.com/google/wire@v0.6.0
```


## 配置使用

> 环境变量 > yaml配置文件 > 默认配置

- 环境变量方法

```shell
export DDD_HTTP_PORT=8080
export DDD_DB_NAME=test.db
```

## grpc

```shell
sudo apt  install protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
```

```shell
go get google.golang.org/grpc@v1.62.2
```
