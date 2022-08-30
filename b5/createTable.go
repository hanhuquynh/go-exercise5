package b4

import (
	"log"
)

func CreateTable() {
	engine := ConnectDB()
	err := engine.CreateTables(new(User))
	if err != nil {
		log.Fatal("Create table error: ", err)
	} else {
		log.Println("Create table successfully!")
	}
}
