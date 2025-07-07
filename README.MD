# swagger-ui

## Golang一键构建
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o swagger-linux-amd64 main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o swagger-windows-amd64.exe main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o swagger-darwin-amd64 main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -trimpath -o swagger-darwin-arm64 main.go
```

## ⚠️ 因为资源是使用的相对目录加载的，所以服务启动的时候，要在项目所在目录启动，不然会会出现not found问题。比如使用systemd进程守护，要使用workdir来说明运行的位置

## 使用

```bash
./swagger-darwin-arm64 -secret 123456 -port 8900
```

### 快速上传本地文件到目标服务
```bash
curl -F "file=@docs/example.test.json" \
    -F "filename=ttt.json" \
    -F "secret=password" \
    http://localhost:8900/upload
```

### 获取文档文件列表

```bash
curl http://localhost:8900/list
```

### 刷新文档
```bash
curl http://localhost:8900/refresh
```

### 简单守护启动

```bash
nohup ./swagger -secret password >> server.log 2>&1 &
```
