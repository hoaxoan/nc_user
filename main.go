package main

import (
	"fmt"
	"log"

	"github.com/hoaxoan/nc_course/nc_user/config"
	db "github.com/hoaxoan/nc_course/nc_user/db"
	md "github.com/hoaxoan/nc_course/nc_user/middleware"
	us "github.com/hoaxoan/nc_course/nc_user/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(md.Logger())
	client, err := db.Connection()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	usRepo := &us.UserRepository{client}
	tokenService := &us.TokenService{usRepo}
	srv := &us.UserService{usRepo, tokenService}
	us.NewUserHandler(e, srv)

	log.Println(e.Start(":9090"))
}
