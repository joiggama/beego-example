package controllers

import(
  "encoding/json"
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

func (c *ProductsController) Post() {
  db := orm.NewOrm()

  var product models.Product
  json.Unmarshal(c.Ctx.Input.RequestBody, &product)

  _, err := db.Insert(&product)

  if err == nil {
    c.Data["json"] = &product
  } else {
    c.Data["json"] = map[string]string{ "error" : err.Error() }
  }
  c.ServeJson()
}
