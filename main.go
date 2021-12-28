package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
)

func main() {
	app := &cli.App{
		Name:        "Swagger UI",
		Usage:       "simple swagger json preview.",
		Description: "预览多个swagger json文件",
		Commands: []*cli.Command{
			{
				Name:  "serve",
				Usage: "启动web服务器",
				Action: func(context *cli.Context) error {
					http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("page"))))
					http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("docs"))))
					err := http.ListenAndServe(":8080", nil)
					if err != nil {
						log.Println(err)
					}
					return nil
				},
			},
			{
				Name:  "init",
				Usage: "初始化",
				Action: func(context *cli.Context) error {
					Init()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
