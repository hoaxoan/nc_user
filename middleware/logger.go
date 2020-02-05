package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
)

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			err := next(c)
			resp := c.Response()
			log.Println(req.URL.Host, req.URL.Path, req.Method, resp.Status)
			//udp metric
			return err
		}
	}
}
