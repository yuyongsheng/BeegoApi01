package main

import (
	_ "BeegoApi01/docs"
	_ "BeegoApi01/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("console", "")
	beego.SetLevel(beego.LevelDebug)
	beego.Debug("Debug ~~~")

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
