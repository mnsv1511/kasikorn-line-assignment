package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/handler/dto"
)

func (h *HandlerImpl) GetUserGreeting(c echo.Context) error {
	// Mock user id
	userId := "1"

	response, err := h.svc.GetUserGreeting(c, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.StatusCode{}.FromDomain(response.Status))
	}

	return c.JSON(http.StatusOK, dto.GetUserGreetingResponse{}.FromDomain(response))
}
