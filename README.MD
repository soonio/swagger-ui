# swagger服务

## 依赖

  - [go-frame](https://goframe.org)

## 安装构建

```
go mod tidy
gf build main.go
mv bin/linux_amd64/swg /var/www/matrix
mv resource/index.html /var/www/matrix
```

## 目录结构
```bash
[root@localhost swagger]# tree
.
├── index.html
├── runtime
│   └── server
│       └── 2021-05-26.log
├── swagger
│   └── swagger.json
└── swg

3 directories, 4 files
```

## 基于supervisor进程守护

```ini
[program:swg]
directory=/var/www/matrix
command=/var/www/matrix/swg
user=root
autostart=true
autorestart=true
startsecs=3
startretries=3
stderr_logfile=/var/log/supervisor/swg-err.log
stdout_logfile=/var/log/supervisor/swg-out.log
```