package http

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hoaxoan/nc_user/config"
	"github.com/hoaxoan/nc_user/model"
	"github.com/hoaxoan/nc_user/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type userHandler struct {
	UUcase user.Usecase
}

func NewUserHandler(e *echo.Echo, uc user.Usecase) {
	handler := &userHandler{
		UUcase: uc,
	}
	PrivateRoute(e, handler)

	PublicRoute(e, handler)
}

func PrivateRoute(e *echo.Echo, handler *userHandler) {
	JWTConfig := middleware.JWTConfig{
		SigningKey: []byte(config.Config.JWTSecret.JWTKey),
		Claims:     &model.CustomClaims{},
	}

	g := e.Group("/api/user/v1/private")
	g.Use(middleware.JWTWithConfig(JWTConfig))
	g.PUT("/user", handler.Update)
}

func PublicRoute(e *echo.Echo, handler *userHandler) {
	g := e.Group("/api/user/v1/public")
	g.PATCH("/user/login", handler.Auth)
	g.POST("/user/register", handler.Register)
}

func (h *userHandler) Register(ctx echo.Context) error {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	var res model.UserResponse
	err := h.UUcase.Create(ctx.Request().Context(), &user, &res)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	return ctx.JSON(http.StatusOK, res.User)
}

func (h *userHandler) Update(ctx echo.Context) error {
	userToken := ctx.Get("user").(*jwt.Token)
	if userToken == nil || userToken.Valid == false {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "token invalid"})
	}
	claims := userToken.Claims.(*model.CustomClaims)

	var resToken model.Token
	reqToken := model.Token{Token: userToken.Raw}
	if err := h.UUcase.ValidateToken(ctx.Request().Context(), &reqToken, &resToken); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	user.Email = claims.User.Email

	var res model.UserResponse
	err := h.UUcase.Update(ctx.Request().Context(), &user, &res)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	return ctx.JSON(http.StatusOK, res.User)
}

func (h *userHandler) Auth(ctx echo.Context) error {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	var token model.Token
	err := h.UUcase.Auth(ctx.Request().Context(), &user, &token)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	return ctx.JSON(http.StatusOK, token)
}
