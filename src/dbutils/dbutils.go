package dbutils

import (
	//"github.com/go-xorm/core"
	//"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func GetDb() *xorm.Engine {
	//log := utils.GetLog("dbutils")
	if engine == nil {
		var err error
		engine, err = xorm.NewEngine("mysql", "root:beta@/demo?charset=utf8")
		//("postgres", "user=test password=test dbname=test sslmode=disable")
		//"postgres", "postgres://postgres:beta@localhost:5432/public"
		if err != nil {
			panic(err.Error())
		}
		engine.Ping()
		//		engine.ShowSQL(true)
		//		engine.Logger().SetLevel(core.LOG_DEBUG)
		//		prefixMapper := core.NewPrefixMapper(core.GonicMapper{}, "beta_")
		//		engine.SetTableMapper(prefixMapper)
	}
	return engine
}
