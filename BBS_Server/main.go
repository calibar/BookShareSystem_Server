package main

import (
"BBS_Server/models"
_ "BBS_Server/routers"
"github.com/astaxie/beego/plugins/cors"

"github.com/astaxie/beego"
"github.com/astaxie/beego/orm"
_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:lmx1993917@tcp(127.0.0.1:3306)/book_borrow_system")
}

func main() {
	go models.CheckExpiration()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	beego.Run()

}

