# broker

## listener

```shell
# arm 模块
GOOS=linux GOARCH=arm64 go build -o broker main.go

# x86-64 模块
GOOS=linux GOARCH=amd64 go build -o broker main.go
```