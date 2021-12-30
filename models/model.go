package models

import (
	"github.com/beego/beego/v2/client/orm"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	//註冊驅動, 原始文件 orm.RegisterDriver("sqlite", orm.DR_Sqlite), 要去掉_
	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	// 設定預設資料庫
	//資料庫存放位置：./datas/test.db ， 資料庫別名：default
	orm.RegisterDataBase("default", "sqlite3", "./datas/test.db")

	//一起註冊
	orm.RegisterModel(new(Meal), new(Groups), new(Record))

	//如果 table 不存在則建立 table，不然跳過
	orm.RunSyncdb("default", false, true)
}
