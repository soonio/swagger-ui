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

  - page目录下的文件来自[swagger-api/swagger-ui](https://github.com/swagger-api/swagger-ui/releases) v4.15.5
  - 项目原理本身是通过修改swagger-ui官方包index.html，注入一个select选框，用于切换多个swagger文档, 具体内容如下
    ```javascript
    let sites = ['ConfigPlaceholder']
      window.onload = function () {

        let prefix = document.location.origin
        sites.map((v) => {
          v.url = prefix + "/" + v.url.replace('./', "")
          return v
        })

        let selectUrl = function (url) {
          window.ui = SwaggerUIBundle({
            url: url,
            dom_id: '#swagger-ui',
            deepLinking: true,
            presets: [
              SwaggerUIBundle.presets.apis,
              SwaggerUIStandalonePreset
            ],
            plugins: [
              SwaggerUIBundle.plugins.DownloadUrl
            ],
            layout: "StandaloneLayout"
          })
        }
        selectUrl(sites[0].url)
        let items = document.getElementById('api-select');
        sites.forEach(site => {
          let option = document.createElement('option')
          option.innerHTML = site.name
          option.value = site.url
          items.appendChild(option)
        });
        items.onchange = function () {
          let valOption = this.options[this.selectedIndex].value
          history.pushState({title: "", url: window.location.url}, '', "?url=" + encodeURIComponent(valOption))
          selectUrl(valOption)
        }
      }
    ```
  - 针对sites数据，ConfigPlaceholder作为位置占用符号，通过python脚本解析docs目录下的swagger json，解析处理乘`<select>`标签的的name和value



## Sponsor

  Thank [JetBrains](https://jb.gg/OpenSourceSupport)