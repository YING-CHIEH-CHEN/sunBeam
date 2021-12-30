package main

import (
	_ "breakfast/routers"

	"github.com/astaxie/beego"
)

func init() {
}

func mod(i, j int) bool {
	return i%j == 0
}

func add(i, j int) int {
	return i + j
}

func main() {
	// 註冊 template 用的函數
	beego.AddFuncMap("mod", mod)
	beego.AddFuncMap("add", add)

	beego.Run()
}
