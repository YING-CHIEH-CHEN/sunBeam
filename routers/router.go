package routers

import (
	"breakfast/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/api",
			beego.NSNamespace("/group",
				beego.NSRouter("/remove/:sn", &controllers.GroupController{}, "*:DeleteGroup"),
				// beego.NSRouter("/mem/:sn", &controllers.GroupController{}, "Get:GroupMem"),
				beego.NSRouter("/mem/:sn", &controllers.GroupController{}, "Post:AddGroupMem"),
			),
			beego.NSNamespace("/record",
				beego.NSRouter("/remove/:sn", &controllers.SummaryController{}, "*:DeleteRecord"),
			),
		)
	// ns := beego.NewNamespace("/group",
	// 	beego.NSRouter("/remove/:sn", &controllers.GroupController{}, "*:DeleteGroup"),
	// 	beego.NSRouter("/mem/:sn", &controllers.GroupController{}, "Get:GroupMem"),
	// 	beego.NSRouter("/mem/:sn", &controllers.GroupController{}, "Post:AddGroupMem"),
	// )
	beego.AddNamespace(ns)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/groups", &controllers.GroupController{})
	beego.Router("/group", &controllers.GroupController{}, "Get:Groups")
	beego.Router("/group/mem/:sn", &controllers.GroupController{}, "Get:GroupMem")
	beego.Router("/group/mem/:sn", &controllers.GroupController{}, "Post:AddGroupMem")
	beego.Router("/group/remove/:sn", &controllers.GroupController{}, "Get:RemoveGroupMem")
	beego.Router("/menu", &controllers.MenuController{})
	beego.Router("/summary", &controllers.SummaryController{})

	// =======================================******========================================
	// # method="post"
	// # 請求方式,加到過濾裡面,驗證post請求
	// var FilterMethod = func(ctx *context.Context) {
	// 	if ctx.Input.IsPost() == false { // post請求
	// 		ctx.Redirect(302, "/")
	// 	}
	// }
	// beego.InsertFilter("/user/ajaxgetuserinfo", beego.BeforeRouter, FilterMethod)

	// user := c.GetString("user")
	// if user == "" {
	// 	c.Data["errmsg"] = "請輸入使用者名稱!"
	// }

	// =======================================******========================================
	// 例子,判斷使用者登陸,等級等
	// 加到路由之前 src/hello/routers/router.go 裡面init
	// var FilterUser = func(ctx *context.Context) {
	// ctx.Request.RequestURI // 路徑
	// 	_, ok := ctx.Input.Session("uid").(int) // 斷言
	// 	if !ok {
	// 		ctx.Redirect(302, "/login/index")
	// 	}
	// }
	// /user/*是user下面的所有路由都要使用此過濾器
	// beego.InsertFilter("/user/*", beego.BeforeRouter, FilterUser)

	// var FilterUser = func(ctx *context.Context) {
	// 	uid := ctx.GetCookie("uid")
	// 	if uid == "" {
	// 		ctx.Redirect(302, "/login/index")
	// 	}
	// }
	// beego.InsertFilter("/user/*", beego.BeforeRouter, FilterUser)

	// =======================================******========================================
	// 儲存Cookie,儲存的值是一個字串,使用型別.(string)斷言
	// c.Ctx.SetCookie("uid", maps[0]["Id"].(string), 60*60*24*365, "/", nil, nil, false) // 主鍵Id首字母大寫,1年
	// 或者儲存Session,不需要斷言,重啟服務就會丟失,記憶體session
	// c.SetSession("uid", uid)

	// 讀取
	// uid := c.Ctx.GetCookie("uid")
	// if uid == "" {
	// 	c.Ctx.Redirect(302, "/login/index")
	// } else {
	// 	c.Ctx.Redirect(302, "/user/index")
	// }

	// uid := c.GetSession("uid")
	// if uid == nil {
	// 	c.Ctx.Redirect(302, "/login/index")
	// } else {
	// 	c.Ctx.Redirect(302, "/user/index")
	// }

}
