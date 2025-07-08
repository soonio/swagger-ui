# swagger-ui

### 快速开始
```bash
npm install

npm start

npm start -- --secret 123456 --port 8900
```

### 快速上传本地文件到目标服务
```bash
curl -F "file=@docs/example.test.json" \
    -F "filename=ttt.json" \
    -F "secret=123456" \
    http://localhost:8900/upload
```

### 基于Docker运行

```bash
docker pull soonio/swagger-hub:1.0
docker run -d --name swagger -p 81:8934 soonio/swagger-hub:1.0
```

### 推送到HUB
```bash

docker build --platform linux/amd64 -t soonio/swagger-hub:1.0 .
docker push soonio/swagger-hub:1.0
```