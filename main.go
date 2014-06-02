package main

import (
	_ "github.com/joiggama/beego-example/routers"
	"github.com/astaxie/beego"
  _ "github.com/lib/pq"
  "github.com/astaxie/beego/orm"
)

func init() {
  orm.RegisterDriver("postgres", orm.DR_Postgres)
  orm.RegisterDataBase("default", "postgres", "dbname=example_app_dev sslmode=disable")
}

func main() {
	beego.Run()
}

