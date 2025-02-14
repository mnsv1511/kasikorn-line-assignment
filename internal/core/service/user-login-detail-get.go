package service

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mnsv1511/kasikorn-line-assignment/constant"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/core/service/domain"
)

func (s *ServiceImpl) GetUserLoginDetail(c echo.Context, userId string) (*domain.GetUserLoginDetailResponse, error) {
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		return &domain.GetUserLoginDetailResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.INVALID_REQUEST_CODE),
				Description: string(constant.INVALID_REQUEST_DES),
			},
		}, err
	}

	userData, err := s.repository.GetUser(intUserId)
	if err != nil {
		return &domain.GetUserLoginDetailResponse{
			Status: &domain.StatusCode{
				Code:        string(constant.USER_NOT_FOUND_CODE),
				Description: string(constant.USER_NOT_FOUND_DES),
			},
		}, err
	}

	return &domain.GetUserLoginDetailResponse{
		UserName:     userData.Name,
		UserImageUrl: userData.ImageURL,
	}, nil
}
