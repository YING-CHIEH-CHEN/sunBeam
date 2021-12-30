package controllers

import (
	"breakfast/datas"
	"breakfast/models"
	"html/template"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

// GroupController operations for Group
type GroupController struct {
	beego.Controller
}

// URLMapping ...
func (c *GroupController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Group
// @Param	body		body 	models.Group	true		"body for Group content"
// @Success 201 {object} models.Group
// @Failure 403 body is empty
// @router / [post]
func (c *GroupController) Post() {
	name := c.GetString("gName")
	dates := c.GetStrings("dates")
	for _, dateStr := range dates {
		Group := &models.Groups{
			Name: name,
			Date: dateStr,
		}
		datas.UpdateGroupRecord(Group)
	}
	c.Redirect("/groups", 303)
}

// GetOne ...
// @Title GetOne
// @Description get Group by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Group
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GroupController) Get() {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1).Format("2006-01-02")
	groups := datas.ResponseGroups()
	if len(groups) > 0 {
		c.Data["groups"] = groups
		c.Data["groupsBool"] = true
	}
	c.Data["firstDay"] = tomorrow
	c.Data["Email"] = "jay2hinet@gmail.com"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML()) // 加入 xsrfdata

	c.TplName = "group.tpl"
}

// Put ...
// @Title Put
// @Description update the Group
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Group	true		"body for Group content"
// @Success 200 {object} models.Group
// @Failure 403 :id is not int
// @router /:id [put]
func (c *GroupController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Group
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *GroupController) Delete() {
}

// 刪除已經登記團體
func (c *GroupController) DeleteGroup() {
	groupSn, _ := c.GetInt(":sn")
	Group := &models.Groups{
		Id: groupSn,
	}
	datas.DeleteGroupRecord(Group)
	c.Redirect("/groups", 303)
}

func (c *GroupController) GroupMem() {
	meals := datas.ResponseMenu("m")
	drinks := datas.ResponseMenu("d")

	Groupid, _ := strconv.Atoi(c.GetString(":sn"))
	records, GroupName := datas.ResponseGroupRecords(Groupid)

	if len(GroupName) > 0 {
		c.Data["groupName"] = GroupName
		if len(records) > 0 {
			c.Data["recordsBool"] = true
			c.Data["records"] = records
		}
		thisDay := time.Now()
		// if thisDay.Hour() > 11 {
		if thisDay.Hour() > 3 {
			thisDay = thisDay.AddDate(0, 0, 1)
		}
		c.Data["toDay"] = thisDay.Format("2006-01-02")
		c.Data["meals"] = meals
		c.Data["drinks"] = drinks
		c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML()) // 加入 xsrfdata
		c.TplName = "groupMem.tpl"
	} else {
		c.Redirect("/", 303)
	}
}

func (c *GroupController) RemoveGroupMem() {
	Groupid, _ := strconv.Atoi(c.GetString(":sn"))
	Records, GroupName := datas.ResponseGroupRecords(Groupid)
	c.Data["groupName"] = GroupName
	if len(Records) > 0 {
		c.Data["records"] = Records
		c.Data["recordsBool"] = true
	}
	thisDay := time.Now()
	// if thisDay.Hour() > 11 {
	if thisDay.Hour() > 3 {
		thisDay = thisDay.AddDate(0, 0, 1)
	}
	c.Data["toDay"] = thisDay.Format("2006-01-02")

	c.TplName = "groupRemoveMem.tpl"
}

func (c *GroupController) AddGroupMem() {
	// name := c.GetString("name") // 一般 form 取值
	// sn := c.GetString(":sn")    // 使用restful 引數方式傳值，取值時變數要加 :變數名
	// sn := c.Ctx.Input.Param(":sn")		// 這樣寫也可以

	Meals := datas.ResponseMealsMap()
	Groupid, _ := strconv.Atoi(c.GetString(":sn"))
	Record := &models.Record{
		Name:    c.GetString("name"),
		Date:    c.GetString("toDay"),
		Main:    "m" + c.GetString("main"),
		Drink:   "d" + c.GetString("drink"),
		Groupid: Groupid,
	}
	datas.UpdateRecord(Record)

	c.Data["sn"] = c.GetString(":sn")
	c.Data["name"] = c.GetString("name")

	c.Data["meal"] = Meals["m"+c.GetString("main")]
	c.Data["drink"] = Meals["d"+c.GetString("drink")]
	c.Data["postBool"] = true
	c.Data["bool"] = true
	c.Redirect("/group/mem/"+c.GetString(":sn"), 303)
}

func (c *GroupController) Groups() {
	groups := datas.ResponseTodayGroups()
	_ = groups
	c.Data["recordsGroupBool"] = true
	c.Data["AllGroups"] = groups
	thisDay := time.Now()
	// if thisDay.Hour() > 11 {
	if thisDay.Hour() > 3 {
		thisDay = thisDay.AddDate(0, 0, 1)
	}
	c.Data["toDay"] = thisDay.Format("2006-01-02")
	c.TplName = "groups.tpl"
}
