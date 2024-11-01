package model

type TokenResponse struct {
	Token string `json:"token"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
