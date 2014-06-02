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

// Index
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

// Show
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

// Create
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

func (c *ProductsController) Update()  {}
func (c *ProductsController) Destroy() {}
