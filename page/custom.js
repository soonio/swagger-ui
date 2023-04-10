window.onload = function () {
  const sites = fetch('./custom.json?t=' + (new Date().valueOf())).then(r => r.json()).then(sites => {
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

    selectUrl(sites[0].url)
  })


  const selectUrl = (url) => {
    url = url.startsWith("./") ? document.location.origin + "/" + url.replace('./', "") : url
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
}
