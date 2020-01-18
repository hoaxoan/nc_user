package main

import (
	"fmt"
	"log"

	"github.com/hoaxoan/nc_user/config"
	"github.com/hoaxoan/nc_user/db"
	md "github.com/hoaxoan/nc_user/middleware"
	_userRepository "github.com/hoaxoan/nc_user/user/repository"
	_userSerice "github.com/hoaxoan/nc_user/user/service"
	_userUsecase "github.com/hoaxoan/nc_user/user/usecase"
	_userHttpDeliver "github.com/hoaxoan/nc_user/user/delivery/http"
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

	usRepo := _userRepository.NewUserRepository(client)
	tokenService := _userSerice.NewTokeService(usRepo)
	uu := _userUsecase.NewUserUsecase(usRepo, tokenService)
	_userHttpDeliver.NewUserHandler(e, uu)

	log.Println(e.Start(":9090"))
}
