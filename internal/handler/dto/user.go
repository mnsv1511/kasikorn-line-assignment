package dto

import "github.com/mnsv1511/kasikorn-line-assignment/internal/core/service/domain"

type GetUserGreetingResponse struct {
	Status       *StatusCode `json:"status,omitempty"`
	UserName     string     `json:"user_name"`
	UserGreeting string     `json:"user_greeting"`
}

type GetUserLoginDetailResponse struct {
	Status       *StatusCode `json:"status,omitempty"`
	UserName     string     `json:"user_name"`
	UserImageUrl string     `json:"user_image_url"`
}

func (d GetUserGreetingResponse) FromDomain(dm *domain.GetUserGreetingResponse) *GetUserGreetingResponse {
	return &GetUserGreetingResponse{
		UserName:     dm.UserName,
		UserGreeting: dm.UserGreeting,
	}
}

func (d GetUserLoginDetailResponse) FromDomain(dm *domain.GetUserLoginDetailResponse) *GetUserLoginDetailResponse {
	return &GetUserLoginDetailResponse{
		UserName:     dm.UserName,
		UserImageUrl: dm.UserImageUrl,
	}
}
