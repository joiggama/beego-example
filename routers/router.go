package routers

import (
	"github.com/astaxie/beego"
	"github.com/joiggama/beego-example/controllers"
)

func init() {
	beego.Router("/products", &controllers.ProductsController{}, "get:Index;post:Create")
	beego.Router("/products/:id:int", &controllers.ProductsController{}, "get:Show;put:Update;delete:Destroy")
}
