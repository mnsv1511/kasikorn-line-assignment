package dto

import "github.com/mnsv1511/kasikorn-line-assignment/internal/core/service/domain"

type StatusCode struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (d StatusCode) FromDomain(dm *domain.StatusCode) *StatusCode {
	return &StatusCode{
		Code:        dm.Code,
		Description: dm.Description,
	}
}
