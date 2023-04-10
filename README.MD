# swagger-ui

## 预览

  [https://swag.iosoon.cn](https://swag.iosoon.cn)

## 优点

- 操作方便🐶
- 可以同时部署多个项目的api文档🚀

## 快速使用

- 根据[模版配置](./conf.nginx)配置nginx
- 把swagger生成的接口文件命名为项目名+环境名+json。如`project.env.json`
- 把`project.env.json`放到docs目录下
- 执行`python3 run.py`即可


## 说明

  - page目录下的文件来自[swagger-api/swagger-ui](https://github.com/swagger-api/swagger-ui/releases) v5.0.0
  - 项目原理本身是通过修改swagger-ui官方包index.html，注入一个select选框，用于切换多个swagger文档
  - 具体方法
    - 引入custom.js
    - 引入custom.css



## Sponsor

  Thank [JetBrains](https://jb.gg/OpenSourceSupport)