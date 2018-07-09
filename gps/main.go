package main

import (
	//	"fmt"
	"GpsServer/gps/models"
	_ "GpsServer/gps/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const ()

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:111@tcp(localhost:3306)/gps")
}

func main() {
	//err := orm.RunSyncdb("default", false, true)
	///if err != nil {
	///	fmt.Println("find err = %+v when beego start run")
	///}
	orm.RegisterModel(new(models.GpsT))
	beego.Run()
}
