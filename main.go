package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"sync"
)

//go:embed static
var static embed.FS

var secret = flag.String("secret", "password", "Secret to authenticate with")
var port = flag.String("port", "8900", "端口")

var locker sync.Mutex

func init() {
	_ = os.Mkdir("docs", 0777)
	_, err := os.Stat("index.html")
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
		var buffer bytes.Buffer
		buffer.Write(part1)
		buffer.Write(part2)
		err = os.WriteFile("index.html", buffer.Bytes(), 0644)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	flag.Parse()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(302, "/swagger/")
	})

	fsys, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}

	e.File("/swagger/", "index.html")
	e.GET("/swagger/*", echo.WrapHandler(http.StripPrefix("/swagger/", http.FileServer(http.FS(fsys)))))
	//e.Static("/swagger/*", "static")
	e.Static("/docs/*", "docs")
	e.POST("/upload", Upload)
	e.GET("/refresh", func(c echo.Context) error {
		refresh(c.Logger())
		return c.Redirect(302, "/swagger/")
	})

	e.GET("list", list)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", *port)))

}

type Swagger struct {
	Info struct {
		Title   string `json:"title"`
		Version string `json:"version"`
	} `json:"info"`
}

func Upload(c echo.Context) error {
	if c.FormValue("secret") != *secret {
		return c.JSON(http.StatusBadRequest, "Invalid secret")
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	src, err := file.Open()
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	defer func() { _ = src.Close() }()

	name := file.Filename
	if nName := c.FormValue("filename"); nName != "" {
		name = nName
	}
	dst, err := os.Create(fmt.Sprintf("docs/%s", name))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		c.Logger().Error(err)
		return err
	}

	go refresh(c.Logger())

	return c.JSON(http.StatusOK, echo.Map{"msg": "ok"})
}

func list(c echo.Context) error {
	var options = make([]map[string]string, 0)

	des, err := os.ReadDir("docs")
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	for _, d := range des {
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".json") {
			if bs, err := os.ReadFile("docs/" + d.Name()); err != nil {
				c.Logger().Error(err)
				continue
			} else {
				var s Swagger
				err = json.Unmarshal(bs, &s)
				if err != nil {
					c.Logger().Error(err)
					continue
				}

				options = append(options, map[string]string{
					"title":   s.Info.Title,
					"version": s.Info.Version,
					"url":     fmt.Sprintf("/docs/%s", d.Name()),
				})
			}
		}
	}
	return c.JSONPretty(http.StatusOK, options, "    ")
}

func refresh(logger echo.Logger) {
	locker.Lock()
	defer locker.Unlock()

	des, err := os.ReadDir("docs")
	if err != nil {
		logger.Error(err)
		return
	}

	var options string

	for _, d := range des {
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".json") {
			if bs, err := os.ReadFile("docs/" + d.Name()); err != nil {
				logger.Error(err)
				continue
			} else {
				var s Swagger
				err = json.Unmarshal(bs, &s)
				if err != nil {
					logger.Error(err)
					continue
				}

				options += fmt.Sprintf("{url: \"%s\", name: \"%s-%s\"},\n                ", fmt.Sprintf("/docs/%s", d.Name()), s.Info.Title, s.Info.Version)
			}
		}
	}

	var buffer bytes.Buffer
	buffer.Write(part1)
	buffer.WriteString(options)
	buffer.Write(part2)

	err = os.WriteFile("index.html", buffer.Bytes(), 0644)
	if err != nil {
		logger.Error(err)
	}
}

var part1 = []byte(`
<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Swagger UI</title>
    <link rel="stylesheet" type="text/css" href="./swagger-ui.css"/>
    <link rel="stylesheet" type="text/css" href="index.css"/>
    <link rel="icon" type="image/png" href="./favicon-32x32.png" sizes="32x32"/>
    <link rel="icon" type="image/png" href="./favicon-16x16.png" sizes="16x16"/>
</head>

<body>
<div id="swagger-ui"></div>
<script src="./swagger-ui-bundle.js" charset="UTF-8"></script>
<script src="./swagger-ui-standalone-preset.js" charset="UTF-8"></script>
<script>
    window.onload = function () {
        window.ui = SwaggerUIBundle({
            urls: [
                `)
var part2 = []byte(`{url: "https://petstore.swagger.io/v2/swagger.json", name: "官方样例"}
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
            layout: "StandaloneLayout",
			validatorUrl: null
        });
    };

</script>
</body>
</html>`)
