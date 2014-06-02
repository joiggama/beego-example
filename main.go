package main

import (
	"os"
	"strconv"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"

	_ "github.com/joiggama/beego-example/routers"

	"github.com/codegangsta/envy/lib"
)

func init() {
	envy.Bootstrap()

	orm.RegisterDriver("postgres", orm.DR_Postgres)
	orm.RegisterDataBase("default", "postgres", os.Getenv("DATABASE_URL"))
}

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.HttpPort = port
	}
	beego.Run()
}
