# swagger-ui

## V5 使用golang，一键构建

### 上传文件
```bash
curl -F "file=@docs/example.test.json" \
    -F "filename=ttt.json" \
    -F "secret=password" \
    http://localhost:8934/upload 
```

### 简单启动

```bash
nohup ./swagger -secret password >> server.log 2>&1 &
```