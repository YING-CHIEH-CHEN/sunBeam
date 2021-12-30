package datas

import (
	"breakfast/models"
	"strconv"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// 依照型態查詢 menu 總表
func ResponseMenu(mealType string) []*models.Meal {
	o := orm.NewOrm()

	var Meals []*models.Meal
	_, _ = o.QueryTable("Meal").Filter("Loc__istartswith", mealType).OrderBy("Loc").All(&Meals)
	return Meals
}

// 更新 menu 紀錄
func UpdateMenu(meals map[string]string) {
	o := orm.NewOrm()

	_, _ = o.QueryTable("meal").Filter("loc", "m1").Update(orm.Params{"name": meals["m1"]})
	_, _ = o.QueryTable("meal").Filter("loc", "m2").Update(orm.Params{"name": meals["m2"]})
	_, _ = o.QueryTable("meal").Filter("loc", "m3").Update(orm.Params{"name": meals["m3"]})
	_, _ = o.QueryTable("meal").Filter("loc", "m4").Update(orm.Params{"name": meals["m4"]})
	_, _ = o.QueryTable("meal").Filter("loc", "m5").Update(orm.Params{"name": meals["m5"]})
	_, _ = o.QueryTable("meal").Filter("loc", "m6").Update(orm.Params{"name": meals["m6"]})
	_, _ = o.QueryTable("meal").Filter("loc", "d1").Update(orm.Params{"name": meals["d1"]})
	_, _ = o.QueryTable("meal").Filter("loc", "d2").Update(orm.Params{"name": meals["d2"]})
	_, _ = o.QueryTable("meal").Filter("loc", "d3").Update(orm.Params{"name": meals["d3"]})
	_, _ = o.QueryTable("meal").Filter("loc", "d4").Update(orm.Params{"name": meals["d4"]})
	_, _ = o.QueryTable("meal").Filter("loc", "d5").Update(orm.Params{"name": meals["d5"]})
	_, _ = o.QueryTable("meal").Filter("loc", "d6").Update(orm.Params{"name": meals["d6"]})
}

// 新增團體登記
func UpdateGroupRecord(record *models.Groups) {
	o := orm.NewOrm()

	// 插入表
	_, _ = o.Insert(record)
}

// 刪除已經登記團體
func DeleteGroupRecord(record *models.Groups) {
	o := orm.NewOrm()

	// 插入表
	_, _ = o.Delete(record)
}

// 已經登記團體總表
func ResponseGroups() []*models.Groups {
	now := time.Now().Format("2006-01-02")

	o := orm.NewOrm()

	var Groups []*models.Groups
	_, _ = o.QueryTable("Groups").Filter("date__gte", now).OrderBy("date", "name").All(&Groups)
	return Groups
}

// 查詢某筆訂餐紀錄
func GetRecord(recordSn int) *models.Record {
	o := orm.NewOrm()
	var Record models.Record
	_ = o.QueryTable("Record").Filter("id", recordSn).One(&Record)
	return &Record
}

// 新增個人訂餐
func UpdateRecord(record *models.Record) {
	o := orm.NewOrm()
	// 插入表
	_, _ = o.Insert(record)
}

// 刪除個人訂餐
func DeleteRecord(record *models.Record) {
	o := orm.NewOrm()

	// 插入表
	_, _ = o.Delete(record)
}

// 今日個人訂餐總表
func ResponseTodayRecords() []*models.Record {
	thisDay := time.Now()
	// if thisDay.Hour() > 11 {
	if thisDay.Hour() > 3 {
		thisDay = thisDay.AddDate(0, 0, 1)
	}
	// return ResponseRecords(now.Format("2006-01-02"))
	return ResponseRecords(thisDay.Format("2006-01-02"))
}

// 某日個人訂餐總表
func ResponseRecords(queryDate string) []*models.Record {
	o := orm.NewOrm()

	var Records []*models.Record
	_, _ = o.QueryTable("Record").Filter("Date", queryDate).Filter("groupid", 0).OrderBy("Name").All(&Records)
	return Records
}

// 今日某團體訂餐總表
func ResponseGroupRecords(queryGroupId int) ([]*models.Record, string) {
	o := orm.NewOrm()

	var Records []*models.Record
	_, _ = o.QueryTable("Record").Filter("groupid", queryGroupId).OrderBy("Name").All(&Records)
	if len(Records) > 0 {
		Meals := ResponseMealsMap()
		for _, r := range Records {
			r.Main = Meals[r.Main]
			r.Drink = Meals[r.Drink]
		}
	}

	var Group models.Groups

	thisDay := time.Now()
	// if thisDay.Hour() > 11 {
	if thisDay.Hour() > 3 {
		thisDay = thisDay.AddDate(0, 0, 1)
	}
	_, _ = o.QueryTable("Groups").Filter("id", queryGroupId).Filter("date", thisDay.Format("2006-01-02")).All(&Group)

	return Records, Group.Name
}

// 今日某團體訂餐總和
func ResponseGroupsSum(queryGroup string) ([6]string, [6]string) {
	o := orm.NewOrm()
	var resMain []OrderMainSum
	var resDrink []OrderMainSum
	ggM := [6]string{"0", "0", "0", "0", "0", "0"}
	ggD := [6]string{"0", "0", "0", "0", "0", "0"}

	thisDay := time.Now()
	// if thisDay.Hour() > 11 {
	if thisDay.Hour() > 3 {
		thisDay = thisDay.AddDate(0, 0, 1)
	}
	_, _ = o.Raw("select r.main,count(*) as res, m.cnt from record as r,meal as m where r.date=? and groupid=? and r.main like 'm%' and r.main=m.loc group by r.main order by r.main", thisDay.Format("2006-01-02"), queryGroup).QueryRows(&resMain)
	for i := 0; i < len(resMain); i++ {
		ggM[resMain[i].Cnt-1] = strconv.Itoa(resMain[i].Res)
	}

	_, _ = o.Raw("select r.main,count(*) as res, m.cnt from record as r,meal as m where r.date=? and groupid=? and r.drink like 'd%' and r.drink=m.loc group by r.drink order by r.drink", thisDay.Format("2006-01-02"), queryGroup).QueryRows(&resDrink)
	for i := 0; i < len(resDrink); i++ {
		ggD[resDrink[i].Cnt-1] = strconv.Itoa(resDrink[i].Res)
	}
	return ggM, ggD
}

// 今日團體
func ResponseTodayGroups() map[int]string {
	GroupsMap := make(map[int]string)
	o := orm.NewOrm()
	var Groups []*models.Groups

	thisDay := time.Now()
	// if thisDay.Hour() > 11 {
	if thisDay.Hour() > 3 {
		thisDay = thisDay.AddDate(0, 0, 1)
	}
	_, _ = o.QueryTable("Groups").Filter("date", thisDay.Format("2006-01-02")).All(&Groups)
	for _, g := range Groups {
		GroupsMap[g.Id] = g.Name
	}
	return GroupsMap
}

type OrderMainSum struct {
	Cnt  int    `json:"Cnt" orm:"column(cnt)"`
	Res  int    `json:"Res" orm:"column(res)"`
	Main string `json:"Main" orm:"column(main)"`
}
type OrderDrinkSum struct {
	Cnt   int    `json:"Cnt" orm:"column(cnt)"`
	Res   int    `json:"Res" orm:"column(res)"`
	Drink string `json:"Drink" orm:"column(drink)"`
}

type GroupDetail struct {
	Id        string
	Name      string
	SumMeals  [6]string
	SumDrinks [6]string
}

type RecordDetail struct {
	Id    string
	Name  string
	Meal  string
	Drink string
}

// 今日全部團體訂餐總表
// []GroupDetail
func ResponseGroupsRecords() map[string]GroupDetail {
	o := orm.NewOrm()

	var groupList orm.ParamsList

	thisDay := time.Now()
	// if thisDay.Hour() > 11 {
	if thisDay.Hour() > 3 {
		thisDay = thisDay.AddDate(0, 0, 1)
	}
	_, _ = o.Raw("select id from groups where date=?", thisDay.Format("2006-01-02")).ValuesFlat(&groupList)

	for index, row := range groupList {
		gg := make([]string, len(groupList))
		gg[index] = row.(string)
	}
	AllGroups := make(map[string]GroupDetail)
	_ = AllGroups

	groupName := ResponseTodayGroups()

	for index, row := range groupList {
		groupId, _ := strconv.Atoi(row.(string))
		SumMeals, SumDrinks := ResponseGroupsSum(row.(string))
		thisGroup := &GroupDetail{
			Id:        row.(string),
			Name:      groupName[groupId],
			SumMeals:  SumMeals,
			SumDrinks: SumDrinks,
		}
		AllGroups[strconv.Itoa(index)] = *thisGroup
	}

	var resMain []OrderMainSum
	var resDrink []OrderDrinkSum

	SumMeals := [6]string{"0", "0", "0", "0", "0", "0"}
	SumDrinks := [6]string{"0", "0", "0", "0", "0", "0"}
	Locs := make(map[string]int)
	Locs["m1"], Locs["d1"] = 0, 0
	Locs["m2"], Locs["d2"] = 1, 1
	Locs["m3"], Locs["d3"] = 2, 2
	Locs["m4"], Locs["d4"] = 3, 3
	Locs["m5"], Locs["d5"] = 4, 4
	Locs["m6"], Locs["d6"] = 5, 5
	_ = SumMeals
	_, _ = o.Raw("select r.main,count(*) as res, m.cnt from record as r,meal as m where r.date='" + thisDay.Format("2006-01-02") + "' and r.main=m.loc group by r.main order by main").QueryRows(&resMain)
	SumMealsCnt := 0
	_ = SumMealsCnt
	for i := 0; i < len(resMain); i++ {
		SumMeals[Locs[resMain[i].Main]] = strconv.Itoa(resMain[i].Res)
	}
	SumDrinksCnt := 0
	_ = SumDrinksCnt
	_, _ = o.Raw("select r.drink,count(*) as res, m.cnt from record as r,meal as m where r.date='" + thisDay.Format("2006-01-02") + "' and r.drink=m.loc group by r.drink order by r.drink").QueryRows(&resDrink)
	for i := 0; i < len(resDrink); i++ {
		SumDrinks[Locs[resDrink[i].Drink]] = strconv.Itoa(resDrink[i].Res)
	}

	return AllGroups
}

// 今日全體訂餐總表
func ResponseAllRecords() ([6]string, [6]string) {
	o := orm.NewOrm()

	var resMain []OrderMainSum
	var resDrink []OrderDrinkSum

	SumMeals := [6]string{"0", "0", "0", "0", "0", "0"}
	SumDrinks := [6]string{"0", "0", "0", "0", "0", "0"}
	Locs := make(map[string]int)
	Locs["m1"], Locs["d1"] = 0, 0
	Locs["m2"], Locs["d2"] = 1, 1
	Locs["m3"], Locs["d3"] = 2, 2
	Locs["m4"], Locs["d4"] = 3, 3
	Locs["m5"], Locs["d5"] = 4, 4
	Locs["m6"], Locs["d6"] = 5, 5
	_ = SumMeals

	thisDay := time.Now()
	// if thisDay.Hour() > 11 {
	if thisDay.Hour() > 3 {
		thisDay = thisDay.AddDate(0, 0, 1)
	}
	_, _ = o.Raw("select r.main,count(*) as res, m.cnt from record as r,meal as m where r.date='" + thisDay.Format("2006-01-02") + "' and r.main=m.loc group by r.main order by main").QueryRows(&resMain)
	for i := 0; i < len(resMain); i++ {
		SumMeals[Locs[resMain[i].Main]] = strconv.Itoa(resMain[i].Res)
	}
	_, _ = o.Raw("select r.drink,count(*) as res, m.cnt from record as r,meal as m where r.date='" + thisDay.Format("2006-01-02") + "' and r.drink=m.loc group by r.drink order by r.drink").QueryRows(&resDrink)
	for i := 0; i < len(resDrink); i++ {
		SumDrinks[Locs[resDrink[i].Drink]] = strconv.Itoa(resDrink[i].Res)
	}

	return SumMeals, SumDrinks
}

// menu 對照表	"m1":"xxxx", "d1":"oooo"
func ResponseMealsMap() map[string]string {
	Meals := make(map[string]string)
	meals := ResponseMenu("m")
	drinks := ResponseMenu("d")

	for _, m := range meals {
		Meals[m.Loc] = m.Name
	}
	for _, d := range drinks {
		Meals[d.Loc] = d.Name
	}
	return Meals
}
