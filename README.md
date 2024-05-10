# broker

本项目主要是读取 mqtt broker 的所有 topic，用来确认是否存在数据传输的问题。

```shell
# arm 模块
GOOS=linux GOARCH=arm64 go build -o broker main.go

# x86-64 模块
GOOS=linux GOARCH=amd64 go build -o broker main.go
```