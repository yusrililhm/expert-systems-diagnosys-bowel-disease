package dto

type AdminLoginPayload struct {
	Username string `json:"Username" form:"Username" valid:"required~Username can't be empty"`
	Password string `json:"password" form:"password" valid:"required~Password can't be empty, length(8|255)~Minimum password is 8 characters"`
}
