package main

import (
	"github.com/astaxie/beego"

	_ "github.com/go-steven/cms/src/routers"
	_ "github.com/go-steven/cms/src/service"
)

func main() {
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogger("console", "")
	beego.Run()
}
