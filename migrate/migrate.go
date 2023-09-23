package main

import (
	"log"
	"recruit-info-service/db"
	"recruit-info-service/model"
)

func main() {
	dbConn := db.NewDB()
	defer log.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(
		&model.User{}, 
		&model.Company{}, 
		&model.Technology{}, 
		&model.CompanyTechnology{},
		&model.TechnologyTag{},
		&model.TechnologyTechnologyTag{},
		&model.Like{},
		&model.Comment{},
	)
}