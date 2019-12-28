package user

import (
	md "github.com/hoaxoan/nc_user/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	UserSrv *UserService
}

func NewUserHandler(e *echo.Echo, srv *UserService) {
	handler := &UserHandler{
		UserSrv: srv,
	}
	g := e.Group("/v1/public/user")
	g.PATCH("/login", handler.Auth)
	g.POST("/register", handler.Register)

	e.PUT("/v1/private/user", handler.Update)
}

func (h *UserHandler) Register(ctx echo.Context) error {
	var user md.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, md.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	var res md.UserResponse
	err := h.UserSrv.Create(ctx.Request().Context(), &user, &res)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, md.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	return ctx.JSON(http.StatusOK, res.User)
}

func (h *UserHandler) Update(ctx echo.Context) error {
	var user md.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, md.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	var res md.UserResponse
	err := h.UserSrv.Update(ctx.Request().Context(), &user, &res)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, md.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	return ctx.JSON(http.StatusOK, res.User)
}

func (h *UserHandler) Auth(ctx echo.Context) error {
	var user md.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, md.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	var token md.Token
	err := h.UserSrv.Auth(ctx.Request().Context(), &user, &token)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, md.Error{Code: http.StatusBadRequest, Description: "bad request"})
	}

	return ctx.JSON(http.StatusOK, token)
}
