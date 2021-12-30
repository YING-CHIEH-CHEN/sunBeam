package controllers

import (
	"breakfast/datas"
	"html/template"

	"github.com/astaxie/beego"
)

// MenuController operations for Menu
type MenuController struct {
	beego.Controller
}

// URLMapping ...
func (c *MenuController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Menu
// @Param	body		body 	models.Menu	true		"body for Menu content"
// @Success 201 {object} models.Menu
// @Failure 403 body is empty
// @router / [post]
func (c *MenuController) Post() {
	Meals := make(map[string]string)
	Meals["m1"] = c.GetString("m1")
	Meals["m2"] = c.GetString("m2")
	Meals["m3"] = c.GetString("m3")
	Meals["m4"] = c.GetString("m4")
	Meals["m5"] = c.GetString("m5")
	Meals["m6"] = c.GetString("m6")
	Meals["d1"] = c.GetString("d1")
	Meals["d2"] = c.GetString("d2")
	Meals["d3"] = c.GetString("d3")
	Meals["d4"] = c.GetString("d4")
	Meals["d5"] = c.GetString("d5")
	Meals["d6"] = c.GetString("d6")
	datas.UpdateMenu(Meals)

	// 303: clients must use a GET request
	// 307: clients must use the original method (eg: POST if a POST was originally used)
	c.Redirect("/menu", 303)
}

// GetOne ...
// @Title GetOne
// @Description get Menu by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Menu
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MenuController) Get() {
	c.Data["meals"] = datas.ResponseMenu("m")
	c.Data["drinks"] = datas.ResponseMenu("d")

	c.Data["Email"] = "jay2hinet@gmail.com"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML()) // 加入 xsrfdata

	c.TplName = "menu.tpl"
}

// Put ...
// @Title Put
// @Description update the Menu
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Menu	true		"body for Menu content"
// @Success 200 {object} models.Menu
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MenuController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Menu
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MenuController) Delete() {

}
