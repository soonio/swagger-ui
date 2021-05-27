# swagger服务

## 依赖

  - [go-frame](https://goframe.org)

## 安装构建

```
go mod tidy
gf build 
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