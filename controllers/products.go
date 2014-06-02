package controllers

import(
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"

  "github.com/joiggama/beego-example/models"
)

type ProductsController struct {
  beego.Controller
}

func (c *ProductsController) Get() {
  db := orm.NewOrm()

  var products []models.Product
  qs := db.QueryTable("products")

  limit, _ := c.GetInt("limit")
  offset, _ := c.GetInt("offset")

  qs.Limit(limit, offset).All(&products)
  c.Data["json"] = &products
  c.ServeJson()
}
