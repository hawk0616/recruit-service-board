package main

import (
	"fmt"
	"recruit-info-service/db"
	"recruit-info-service/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Company{}, &model.Technology{}, &model.CompanyTechnology{})
}