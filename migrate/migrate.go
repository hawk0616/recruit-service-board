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
	if err := dbConn.AutoMigrate(
		&model.User{}, 
		&model.Company{}, 
		&model.Technology{}, 
		&model.CompanyTechnology{},
		&model.TechnologyTag{},
		&model.TechnologyTechnologyTag{},
		&model.Like{},
		&model.Comment{},
	); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
}