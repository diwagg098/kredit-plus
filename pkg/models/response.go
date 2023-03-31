package models

type Responseservice struct {
	Status       int         `json:"status"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
	Pagination   interface{} `json:"pagination"`
	Message      string      `json:"message"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	User        User
}
