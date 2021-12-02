# swagger-ui

## 优点

- 操作方便🐶
- 可以同时部署多个项目的api文档🚀

## 快速使用

- 配置nginx
- 把swagger生成的接口文件命名为项目名+环境名+json。如`project.env.json`
- 把`project.env.json`放到docs目录下
- 执行`./swag-ui`即可

## 自定义

- 修改main文件
- 重新编译swag-ui命令

  ```
  go build -o swag-ui main.go
  ```

## 说明

  - page目录下的文件来自[swagger-api/swagger-ui](https://github.com/swagger-api/swagger-ui/releases) v4
