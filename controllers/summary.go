package controllers

import (
	"breakfast/datas"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

// SummaryController operations for Summary
type SummaryController struct {
	beego.Controller
}

// URLMapping ...
func (c *SummaryController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Summary
// @Param	body		body 	models.Summary	true		"body for Summary content"
// @Success 201 {object} models.Summary
// @Failure 403 body is empty
// @router / [post]
func (c *SummaryController) Post() {
}

// GetOne ...
// @Title GetOne
// @Description get Summary by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Summary
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SummaryController) Get() {
	// 主食 menu : []string
	c.Data["meals"] = datas.ResponseMenu("m")
	// 飲料 menu : []string
	c.Data["drinks"] = datas.ResponseMenu("d")
	// 主食小計, 飲料小計 : []string, []string
	c.Data["mealsAll"], c.Data["drinksAll"] = datas.ResponseAllRecords()

	mealsAll, drinksAll := datas.ResponseAllRecords()
	tmpMealSum, tmpDrinkSum := 0, 0
	tmp := 0
	for i := 0; i < 6; i++ {
		tmp, _ = strconv.Atoi(mealsAll[i])
		tmpMealSum += tmp
		tmp, _ = strconv.Atoi(drinksAll[i])
		tmpDrinkSum += tmp
	}
	// 主食總數 : int
	c.Data["MealSum"] = tmpMealSum
	// 飲料總數 : int
	c.Data["DrinkSum"] = tmpDrinkSum

	c.Data["recordsGroupBool"] = true
	// 全部的團體訂餐資料 : [] data.GroupDetail
	// {Id : string, Name : string, SumMeals : [6]string, SumDrinks : [6]string}
	AllGroups := datas.ResponseGroupsRecords()
	if len(AllGroups) > 0 {
		c.Data["AllGroups"] = AllGroups
	}

	// menu 名稱對照, map[string]string
	// map["m1"] : "巧克力吐司"
	Meals := datas.ResponseMealsMap()

	Records := datas.ResponseTodayRecords()
	if len(Records) > 0 {
		for _, r := range Records {
			r.Main = Meals[r.Main]
			r.Drink = Meals[r.Drink]
		}
		// 個人訂餐總表, []data.RecordDetail
		// Id : string, Name : string, Meal : string, Drink : string
		c.Data["records"] = Records
		c.Data["recordsBool"] = true
	}

	thisDay := time.Now()
	// if thisDay.Hour() > 11 {
	if thisDay.Hour() > 3 {
		thisDay = thisDay.AddDate(0, 0, 1)
	}
	c.Data["toDay"] = thisDay.Format("2006-01-02")
	c.Data["toHour"] = thisDay.Format("2006-01-02 03-04")

	c.TplName = "summary.tpl"
}

// Put ...
// @Title Put
// @Description update the Summary
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Summary	true		"body for Summary content"
// @Success 200 {object} models.Summary
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SummaryController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Summary
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SummaryController) Delete() {

}

func (c *SummaryController) DeleteRecord() {
	recordSn, _ := c.GetInt(":sn")

	Record := datas.GetRecord(recordSn)

	datas.DeleteRecord(Record)
	if Record.Groupid == 0 {
		c.Redirect("/", 303)
	} else {
		c.Redirect("/group/mem/"+strconv.Itoa(Record.Groupid), 303)
	}
}
