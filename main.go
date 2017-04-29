package main

import (
	"github.com/astaxie/beego"
	"shopktl/models"
	_ "shopktl/routers"
	"os"
)


func init() {
	models.Connect()
	initArgs()
}

func main() {
	beego.Run()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncbd(false)
			os.Exit(0)
		}
		if v == "-syncdb-force" {
			models.Syncbd(true)
			os.Exit(0)
		}
	}
}
