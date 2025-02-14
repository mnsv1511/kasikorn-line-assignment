package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/core/service"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/handler"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/repository"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	e := echo.New()

	// Init repository
	repo, err := repository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}
	// Init service
	service, err := service.NewService(repo)
	if err != nil {
		log.Fatal(err)
	}
	// Init Handler
	handler, err := handler.NewHadler(service)
	if err != nil {
		log.Fatal(err)
	}

	// Router
	e.GET("/account", handler.GetAccountList)
	e.GET("/account/main", handler.GetAccountMain)
	e.GET("/debit-card/example", handler.GetDebitCardExample)
	e.GET("/transaction", handler.GetTransaction)
	e.GET("/user/greeting", handler.GetUserGreeting)
	e.GET("/user/login", handler.GetUserLoginDetail)

	e.Start(":" + os.Getenv("PORT"))
}
