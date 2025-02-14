package service

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mnsv1511/kasikorn-line-assignment/constant"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/core/service/domain"
)

func (s *ServiceImpl) GetUserGreeting(c echo.Context, userId string) (*domain.GetUserGreetingResponse, error) {
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		return &domain.GetUserGreetingResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.INVALID_REQUEST_CODE),
				Description: string(constant.INVALID_REQUEST_DES),
			},
		}, err
	}

	userData, err := s.repository.GetUser(intUserId)
	if err != nil {
		return &domain.GetUserGreetingResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.USER_NOT_FOUND_CODE),
				Description: string(constant.USER_NOT_FOUND_DES),
			},
		}, err
	}

	userGreetingData, err := s.repository.GetUserGreeting(intUserId)
	if err != nil {
		return &domain.GetUserGreetingResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.USER_GREETING_NOT_FOUND_CODE),
				Description: string(constant.USER_GREETING_NOT_FOUND_DES),
			},
		}, err
	}
	return &domain.GetUserGreetingResponse{
		UserName: userData.Name,
		UserGreeting: userGreetingData.Greeting,
	}, nil
}
