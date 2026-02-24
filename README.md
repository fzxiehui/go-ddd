# ddd 项目构建

## 依赖关系限制

| 层 | 可以依赖 | 不能依赖 | 说明 |
| :------ | :---- | :---- | :------- |
| `domain` | `标准库` | `application` / `infra` / `interface` | 最纯粹业务规则 |
| `application` | `domain` | `interface` | 用例编排层 |
| `infra` | `domain` / `application（事件回调场景允许）` | `interface`  | 技术实现层 |
| `interface` | `application` / `domain(dto)` | `infra` | 对外适配层 |
| `cmd/server`  | 所有层 | 无限制 | 装配层 |

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

- 测试

```shell
# login
grpcurl -plaintext \
  -d '{"username":"root","password":"root"}' \
  localhost:9090 user.v1.AuthService/Login

# get me
grpcurl -plaintext \
-H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMzhlZTczMTQtN2ExMC00ZTZjLWI1NDMtYzUzMjhkMzUxYzk1IiwiZXhwIjoxNzcwOTY1OTc5LCJpYXQiOjE3NzA4Nzk1Nzl9.tmUXu-du04pFrjQ29L1mMs3gPRi40U5MuNZXwPJc8mQ" \
-d '{}' \
localhost:9090 user.v1.AuthService/Me
```
