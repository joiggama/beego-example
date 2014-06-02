package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/joiggama/beego-example/models"
)

type ProductsController struct {
	beego.Controller
}

func (c *ProductsController) Index() {
	db := orm.NewOrm()

	var products []models.Product
	qs := db.QueryTable("products")

	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	qs.Limit(limit, offset).All(&products)
	c.Data["json"] = &products
	c.ServeJson()
}

func (c *ProductsController) Show() {
	db := orm.NewOrm()
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))

	if err != nil {
		panic(err)
	}

	product := models.Product{Id: id}

	err = db.Read(&product)

	if err == nil {
		c.Data["json"] = &product
	} else {
		c.Data["json"] = map[string]string{"error": err.Error()}
	}
	c.ServeJson()
}

func (c *ProductsController) Create() {
	db := orm.NewOrm()

	var product models.Product
	json.Unmarshal(c.Ctx.Input.RequestBody, &product)

	_, err := db.Insert(&product)

	if err == nil {
		c.Data["json"] = &product
	} else {
		c.Data["json"] = map[string]string{"error": err.Error()}
	}
	c.ServeJson()
}

func (c *ProductsController) Update()  {
	db := orm.NewOrm()

	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))

  if err != nil {
    panic(err)
  }

  product := models.Product{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &product)
	_, err = db.Update(&product)

  if err == nil {
    c.Data["json"] = &product
  } else {
    c.Data["json"] = map[string]string{ "error" : err.Error() }
  }
  c.ServeJson()
}

func (c *ProductsController) Destroy() {
	db := orm.NewOrm()

	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))

  if err != nil {
    panic(err)
  }

  _, err = db.Delete(&models.Product{Id: id})

  if err == nil {
    c.Redirect("/products", 204)
  } else {
    c.Data["json"] = map[string]string{ "error" : err.Error() }
    c.ServeJson()
  }
}

func (c *ProductsController) BulkCreate() {
  db := orm.NewOrm()

  var products []models.Product

  json.Unmarshal(c.Ctx.Input.RequestBody, &products)

  _, err := db.InsertMulti(2, &products)

  if err == nil {
    c.Data["json"] = &products
  } else {
    c.Data["json"] = map[string]string{ "error" : err.Error() }
  }

  c.ServeJson()
}
