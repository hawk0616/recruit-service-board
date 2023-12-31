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
	commentValidator := validator.NewCommentValidator()
	technologyValidator := validator.NewTechnologyValidator()

	// repository
	userRepository := repository.NewUserRepository(db)
	companyRepository := repository.NewCompanyRepository(db)
	likeRepository := repository.NewLikeRepository(db)
	commentRepository := repository.NewCommentRepository(db)
	technologyRepository := repository.NewTechnologyRepository(db)
	companyTechnologyRepository := repository.NewCompanyTechnologyRepository(db)
	technologyTagRepository := repository.NewTechnologyTagRepository(db)
	technologyTechnologyTagRepository := repository.NewTechnologyTechnologyTagRepository(db)

	// usecase
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	companyUsecase := usecase.NewCompanyUsecase(companyRepository, companyValidator)
	likeUsecase := usecase.NewLikeUsecase(likeRepository)
	commentUsecase := usecase.NewCommentUsecase(commentRepository, commentValidator)
	technologyUsecase := usecase.NewTechnologyUsecase(technologyRepository, technologyValidator)
	companyTechnologyUsecase := usecase.NewCompanyTechnologyUsecase(companyTechnologyRepository)
	technologyTagUsecase := usecase.NewTechnologyTagUsecase(technologyTagRepository)
	technologyTechnologyTagUsecase := usecase.NewTechnologyTechnologyTagUsecase(technologyTechnologyTagRepository)

	// controller
	userController := controller.NewUserController(userUsecase)
	companyController := controller.NewCompanyController(companyUsecase)
	likeController := controller.NewLikeController(likeUsecase)
	commentController := controller.NewCommentController(commentUsecase)
	technologyController := controller.NewTechnologyController(technologyUsecase)
	companyTechnologyController := controller.NewCompanyTechnologyController(companyTechnologyUsecase)
	technologyTagController := controller.NewTechnologyTagController(technologyTagUsecase)
	technologyTechnologyTagController := controller.NewTechnologyTechnologyTagController(technologyTechnologyTagUsecase)

	// router
	e := router.NewRouter(
		userController,
		companyController,
		likeController,
		commentController,
		technologyController,
		companyTechnologyController,
		technologyTagController,
		technologyTechnologyTagController,
	)

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}