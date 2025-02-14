package domain

type GetUserGreetingResponse struct {
	Status       *StatusCode `json:"status"`
	UserName     string     `json:"user_name"`
	UserGreeting string     `json:"user_greeting"`
}

type GetUserLoginDetailResponse struct {
	Status       *StatusCode `json:"status"`
	UserName     string     `json:"user_name"`
	UserImageUrl string     `json:"user_image_url"`
}
