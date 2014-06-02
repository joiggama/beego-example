package main

import (
  "os"

	"github.com/astaxie/beego"

  _ "github.com/lib/pq"
  "github.com/astaxie/beego/orm"

	_ "github.com/joiggama/beego-example/routers"

  "github.com/codegangsta/envy/lib"
)

func init() {
  envy.Bootstrap()

  orm.RegisterDriver("postgres", orm.DR_Postgres)
  orm.RegisterDataBase("default", "postgres", os.Getenv("DATABASE_URL"))
}

func main() {
	beego.Run()
}

