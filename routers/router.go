package routers

import (
	"github.com/joiggama/beego-example/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/products", &controllers.ProductsController{})
}
