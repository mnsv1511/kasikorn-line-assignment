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

func TestGetUserGreeting(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		mockRepo MockRepository
		want     *domain.GetUserGreetingResponse
		wantErr  error
	}{
		{
			name:  "success get user greeting",
			input: "1",
			mockRepo: MockRepository{
				MockGetUser: func(userId int) (*ent.Users, error) {
					return &ent.Users{
						ID:       1,
						Name:     "Clare",
						ImageURL: "clare_img.png",
					}, nil
				},
				MockGetUserGreeting: func(userId int) (*ent.UserGreetings, error) {
					return &ent.UserGreetings{
						ID:       1,
						UserID:   1,
						Greeting: "Have a nice day",
					}, nil
				},
			},
			want: &domain.GetUserGreetingResponse{
				UserName:     "Clare",
				UserGreeting: "Have a nice day",
			},
		},
		{
			name:  "error req get user greeting",
			input: "a",
			want: &domain.GetUserGreetingResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.INVALID_REQUEST_CODE),
					Description: string(constant.INVALID_REQUEST_DES),
				},
			},
			wantErr: errors.New("error req get user greeting"),
		},
		{
			name:  "error repo get user",
			input: "2",
			mockRepo: MockRepository{
				MockGetUser: func(userId int) (*ent.Users, error) {
					return nil, errors.New("user not found")
				},
			},
			want: &domain.GetUserGreetingResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.USER_NOT_FOUND_CODE),
					Description: string(constant.USER_NOT_FOUND_DES),
				},
			},
			wantErr: errors.New("user not found"),
		},
		{
			name:  "error repo get user greeting",
			input: "1",
			mockRepo: MockRepository{
				MockGetUser: func(userId int) (*ent.Users, error) {
					return &ent.Users{
						ID:       1,
						Name:     "Clare",
						ImageURL: "clare_img.png",
					}, nil
				},
				MockGetUserGreeting: func(userId int) (*ent.UserGreetings, error) {
					return nil, errors.New("user greeting not found")
				},
			},
			want: &domain.GetUserGreetingResponse{
				Status: &domain.StatusCode{
					Code:        string(constant.USER_GREETING_NOT_FOUND_CODE),
					Description: string(constant.USER_GREETING_NOT_FOUND_DES),
				},
			},
			wantErr: errors.New("user greeting not found"),
		},
	}
	for _, testCase := range testCases {
		c := echo.New().AcquireContext()
		t.Run(testCase.name, func(t *testing.T) {
			s := MockService(&testCase.mockRepo)
			got, err := s.GetUserGreeting(c, testCase.input)
			if err != nil {
				assert.Equal(t, testCase.want, got)
			} else {
				assert.Equal(t, testCase.wantErr, err)
			}
		})
	}
}
