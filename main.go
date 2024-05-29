package main

import (
	"Cart/database"
	"Cart/handler"
	"Cart/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	handler.InitValidator()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database.Connect()

	e.POST("/users/register", handler.RegisterUser)
	e.POST("/users/login", handler.LoginUser)

	e.POST("/cart", handler.CreateCart, middlewares.JWTMiddleware)
	e.GET("/cart", handler.GetAllCarts, middlewares.JWTMiddleware)
	e.GET("/cart/:id", handler.GetCart, middlewares.JWTMiddleware)
	e.PATCH("/cart/:id", handler.UpdateCart, middlewares.JWTMiddleware)
	e.DELETE("/cart/:id", handler.DeleteCart, middlewares.JWTMiddleware)

	e.Logger.Fatal(e.Start(":8080"))
}
