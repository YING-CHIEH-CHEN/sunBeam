package controllers

import (
	"breakfast/datas"
	"breakfast/models"
	"html/template"

	"time"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	thisDay := time.Now()
	// if thisDay.Hour() > 11 {
	if thisDay.Hour() > 3 {
		thisDay = thisDay.AddDate(0, 0, 1)
	}

	Meals := datas.ResponseMealsMap()
	meals := datas.ResponseMenu("m")
	drinks := datas.ResponseMenu("d")

	Records := datas.ResponseTodayRecords()
	if len(Records) > 0 {
		for _, r := range Records {
			r.Main = Meals[r.Main]
			r.Drink = Meals[r.Drink]
		}
		c.Data["records"] = Records
		c.Data["recordsBool"] = true
	}
	c.Data["meals"] = meals
	c.Data["drinks"] = drinks
	c.Data["toDay"] = thisDay.Format("2006-01-02")
	c.Data["Email"] = "jay2hinet@gmail.com"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML()) // 加入 xsrfdata
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	// now := time.Now()
	// const (
	// 	stdLongMonth      = "January"
	// 	stdMonth          		= "Jan"
	// 	stdNumMonth		= "1"
	// 	stdZeroMonth		= "01"
	// 	stdLongWeekDay    = "Monday"
	// 	stdWeekDay        	= "Mon"
	// 	stdDay            			= "2"
	// 	stdUnderDay       	= "_2"
	// 	stdZeroDay        		= "02"
	// 	stdHour           			= "15"
	// 	stdHour12         		= "3"
	// 	stdZeroHour12     = "03"
	// 	stdMinute         		= "4"
	// 	stdZeroMinute     = "04"
	// 	stdSecond         	= "5"
	// 	stdZeroSecond     = "05"
	// 	stdLongYear       	= "2006"
	// 	stdYear           		= "06"
	// 	stdPM             		= "PM"
	// 	stdpm             		= "pm"
	// 	stdTZ             		= "MST"
	// 	stdISO8601TZ      = "Z0700"  // prints Z for UTC
	// 	stdISO8601ColonTZ = "Z07:00" // prints Z for UTC
	// 	stdNumTZ          = "-0700"  // always numeric
	// 	stdNumShortTZ     = "-07"    // always numeric
	// 	stdNumColonTZ     = "-07:00" // always numeric
	// 	stdFracSecond0    = ".0", ".00" // trailing zeros included
	// 	stdFracSecond9    = ".9", ".99" // trailing zeros omitted
	// )

	m, d := c.GetString("main"), c.GetString("drink")
	if len(m) < 1 {
		m = ""
	} else {
		m = "m" + m
	}
	if len(d) < 1 {
		d = ""
	} else {
		d = "d" + d
	}
	Record := &models.Record{
		Name:    c.GetString("name"), // 這邊取值很像是用表單的 id
		Date:    c.GetString("toDay"),
		Main:    m,
		Drink:   d,
		Groupid: 0,
	}
	datas.UpdateRecord(Record)

	// 303: clients must use a GET request
	// 307: clients must use the original method (eg: POST if a POST was originally used)
	c.Redirect("/", 303)
}
