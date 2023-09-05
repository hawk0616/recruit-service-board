package main

import (
	"recruit-info-service/controller"
	"recruit-info-service/db"
	"recruit-info-service/repository"
	"recruit-info-service/router"
	"recruit-info-service/usecase"
	"recruit-info-service/validator"
)

func main() {
	// db
	db := db.NewDB()

	// validator
	userValidator := validator.NewUserValidator()
	companyValidator := validator.NewCompanyValidator()

	// repository
	userRepository := repository.NewUserRepository(db)
	companyRepository := repository.NewCompanyRepository(db)
	technologyRepository := repository.NewTechnologyRepository(db)
	companyTechnologyRepository := repository.NewCompanyTechnologyRepository(db)

	// usecase
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	companyUsecase := usecase.NewCompanyUsecase(companyRepository, companyValidator)
	technologyUsecase := usecase.NewTechnologyUsecase(technologyRepository)
	companyTechnologyUsecase := usecase.NewCompanyTechnologyUsecase(companyTechnologyRepository)

	// controller
	userController := controller.NewUserController(userUsecase)
	companyController := controller.NewCompanyController(companyUsecase)
	technologyController := controller.NewTechnologyController(technologyUsecase)
	companyTechnologyController := controller.NewCompanyTechnologyController(companyTechnologyUsecase)

	// router
	e := router.NewRouter(userController, companyController, technologyController, companyTechnologyController)

	// start server
	e.Logger.Fatal(e.Start(":8090"))
}