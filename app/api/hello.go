package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
)

var Hello = helloApi{}

type helloApi struct{}

func (*helloApi) Index(r *ghttp.Request) {

	res := gfile.GetBytes("./index.html")

	r.Response.Write(res)
}
