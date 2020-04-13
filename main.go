package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com:amila-ku/newspal/pkg/handler"
)



func main() {

	
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	// e.POST("/article", createArticle)
	e.GET("/allarticles", GetAllArticles)
	// e.GET("/article/:id", getArticle)
	// e.PUT("/article/:id", updateArticle)
	//e.DELETE("/article/:id", deleteArticle)

	// Static
	//e.Static("/assets", "assets")

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}