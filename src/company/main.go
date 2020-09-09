package main

import (
	_ "company/routers"

	"github.com/astaxie/beego"
)

func index(in int) (out int) {
	out = in + 1
	return out
}

func main() {
	beego.AddFuncMap("index", index)
	beego.Run()
}
