package entities

import (
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

func init() {

	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
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
	// engine.ShowSQL(true)
	// engine.Logger().SetLevel(core.LOG_DEBUG)
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


