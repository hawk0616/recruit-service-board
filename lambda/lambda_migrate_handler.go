package main

import (
	"context"
	"log"
	"recruit-info-service/db"
	"recruit-info-service/model"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (string, error) {
	log.Println("Starting the migration for the recruit info service DB...")
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
	return "Migration completed successfully!", nil
}

func main() {
	lambda.Start(HandleRequest)
}