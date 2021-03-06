package main

import (
	_ "mpEnglishBe/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/mp_english_db?charset=utf8")
}

func main() {
	beego.Run()
}
