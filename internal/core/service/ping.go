package service

import "github.com/labstack/echo/v4"

func (_ *ServiceImpl) TestPing(c echo.Context) (string, error) {
	return "ping", nil
}
