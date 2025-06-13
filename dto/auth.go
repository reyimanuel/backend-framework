package dto

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	StatusCode int           `json:"status"`
	Message    string        `json:"message"`
	Data       TokenResponse `json:"data"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	StatusCode  int    `json:"status"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

type GoogleAuthResponse struct {
	StatusCode int            `json:"status"`
	Message    string         `json:"message"`
	Data       map[string]any `json:"data"`
}
