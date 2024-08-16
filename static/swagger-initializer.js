window.onload = function() {
  console.log("xxx222")
  //<editor-fold desc="Changeable Configuration Block">
  console.log("xxx")

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    // url: "https://petstore.swagger.io/v2/swagger.json",
    urls: [
      {
        url: "https://aaa.com/v2/swagger.json",
        name: "测试"
      },
      {
        url: "https://bbb.com/v2/swagger.json",
        name: "本地"
      }
    ],
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
  });

  //</editor-fold>
};
