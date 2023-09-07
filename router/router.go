package router

import (
	"net/http"
	"os"
	"recruit-info-service/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	uc controller.IUserController, 
	cc controller.ICompanyController,
	lc controller.ILikeController,
	tc controller.ITechnologyController,
	ctc controller.ICompanyTechnologyController,
	ttc controller.ITechnologyTagController,
	tttc controller.ITechnologyTechnologyTagController,
	)*echo.Echo {
	e := echo.New()

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3004", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	// CSRF middleware
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		// CookieSameSite: http.SameSiteNoneMode,
		CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge:   60,
	}))

	// User
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	// Like
	l := e.Group("/users/:userId/likes")
	l.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	l.POST("", lc.CreateLike)
	l.DELETE("/:companyId", lc.DeleteLike)

	// Company
	c := e.Group("/companies")
	c.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	c.GET("", cc.GetAllCompanies)
	c.GET("/:companyId", cc.GetCompanyById)
	c.POST("", cc.CreateCompany)
	c.PUT("/:companyId", cc.UpdateCompany)
	c.DELETE("/:companyId", cc.DeleteCompany)

	// CompanyTechnology
	c.POST("/:companyId/company_technologies", ctc.CreateCompanyTechnology)
	c.DELETE("/:companyId/company_technologies/:id", ctc.DeleteCompanyTechnology)

	// Technology
	t := e.Group("/technologies")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTechnologies)
	t.GET("/:technologyId", tc.GetTechnologyById)
	t.POST("", tc.CreateTechnology)
	t.PUT("/:technologyId", tc.UpdateTechnology)
	t.DELETE("/:technologyId", tc.DeleteTechnology)

	// TechnologyTechnologyTag
	t.POST("/:technologyId/technology_technology_tags", tttc.CreateTechnologyTechnologyTag)
	t.DELETE("/:technologyId/technology_technology_tags/:id", tttc.CreateTechnologyTechnologyTag)

	// TechnologyTag
	tt := e.Group("/technology_tags")
	tt.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	tt.GET("", ttc.GetAllTechnologyTags)
	tt.GET("/:technologyTagId", ttc.GetTechnologyTagById)
	tt.POST("", ttc.CreateTechnologyTag)
	tt.PUT("/:technologyTagId", ttc.UpdateTechnologyTag)
	tt.DELETE("/:technologyTagId", ttc.DeleteTechnologyTag)

	return e
}