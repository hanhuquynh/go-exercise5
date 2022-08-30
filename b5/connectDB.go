package b5

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func ConnectDB() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:hanhuquynh@tcp(127.0.0.1:3306)/ex5?charset=utf8")

	if err != nil {
		log.Fatal("Connect database failed: ", err)
	}

	return engine
}
