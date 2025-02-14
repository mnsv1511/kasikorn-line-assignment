package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/handler/dto"
)

func (h *HandlerImpl) GetAccountMain(c echo.Context) error {
	// Mock user id
	userId := "1"

	response, err := h.svc.GetAccountMain(c, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.StatusCode{}.FromDomain(response.Status))
	}

	return c.JSON(http.StatusOK, dto.GetAccountMainResponse{}.FromDomain(response))
}
