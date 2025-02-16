package service

import (
	"errors"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mnsv1511/kasikorn-line-assignment/constant"
	"github.com/mnsv1511/kasikorn-line-assignment/ent"
	"github.com/mnsv1511/kasikorn-line-assignment/internal/core/service/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetUserLoginDetail(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		mockRepo MockRepository
		want     *domain.GetUserLoginDetailResponse
		wantErr  error
	}{
		{
			name:  "success get user log in detail",
			input: "1",
			mockRepo: MockRepository{
				MockGetUser: func(userId int) (*ent.Users, error) {
					return &ent.Users{
						ID:       1,
						Name:     "Clare",
						ImageURL: "clare_img.png",
					}, nil
				},
			},
			want: &domain.GetUserLoginDetailResponse{
				UserName:     "Clare",
				UserImageUrl: "clare_img.png",
			},
			wantErr: nil,
		},
		{
			name:  "error req get user log in detail",
			input: "a",
			want: &domain.GetUserLoginDetailResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.INVALID_REQUEST_CODE),
					Description: string(constant.INVALID_REQUEST_DES),
				},
			},
			wantErr: errors.New("error req get user log in detail"),
		},
		{
			name:  "error repo get user",
			input: "2",
			mockRepo: MockRepository{
				MockGetUser: func(userId int) (*ent.Users, error) {
					return nil, errors.New("user not found")
				},
			},
			want: &domain.GetUserLoginDetailResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.USER_NOT_FOUND_CODE),
					Description: string(constant.USER_NOT_FOUND_DES),
				},
			},
			wantErr: errors.New("user not found"),
		},
	}
	for _, testCase := range testCases {
		c := echo.New().AcquireContext()
		t.Run(testCase.name, func(t *testing.T) {
			s := MockService(&testCase.mockRepo)
			got, err := s.GetUserLoginDetail(c, testCase.input)
			if err != nil {
				assert.Equal(t, testCase.want, got)
			} else {
				assert.Equal(t, testCase.wantErr, err)
			}
		})
	}
}