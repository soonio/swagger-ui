package main

import (
	_ "swagger/boot"
	_ "swagger/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
