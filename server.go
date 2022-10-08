package main

import (
	"clean_architecture/config"
	"clean_architecture/feature/user/data"
	"clean_architecture/feature/user/delivery"
	"clean_architecture/feature/user/services"
	"clean_architecture/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	//pemanggialn echo
	e := echo.New()

	//pemanggilan config
	cfg := config.NewConfig()
	db := database.InitDB(cfg)

	mdl := data.New(db)
	uServices := services.New(mdl)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	delivery.New(e, &uServices)

	log.Fatal(e.Start(":8000"))

}
