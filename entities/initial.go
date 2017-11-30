package entities

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err= xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	// userInfoTable := new(UserInfo)
	// isExist, _ := engine.IsTableExist(userInfoTable)
	// if(isExist) {
	// 	engine.CreateTables(userInfoTable)
	// }
	engine.SetMapper(core.SameMapper{})
	engine.Sync2(new(UserInfo))
	// engine.ShowWarn = true
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


