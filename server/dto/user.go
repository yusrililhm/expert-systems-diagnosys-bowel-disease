package dto

import "time"

type UserRegisterPayload struct {
	Username  string    `json:"username" form:"username" valid:"required~Username can't be empty"`
	FullName  string    `json:"full_name" form:"full_name" valid:"required~Full name can't be empty"`
	Phone     string    `json:"phone" form:"phone" valid:"required~Phone can't be empty"`
	BirthDate time.Time `json:"birth_date" form:"birth_date" valid:"required~Birth date can't be empty"`
	Gender    bool      `json:"gender" form:"gender"`
	Password  string    `json:"password" form:"password" valid:"required~Password can't be empty, length(8|255)~Minimum password is 8 characters"`
}

type UserLoginPayload struct {
	Phone    string `json:"phone" form:"phone" valid:"required~Phone can't be empty"`
	Password string `json:"password" form:"password" valid:"required~Password can't be empty, length(8|255)~Minimum password is 8 characters"`
}

type ChangePasswordPayload struct {
	OldPassword        string `json:"old_password" form:"old_password" valid:"required~Old password can't be empty, length(8|255)~Minimum password is 8 characters"`
	NewPassword        string `json:"new_password" form:"new_password" valid:"required~New password can't be empty, length(8|255)~Minimum password is 8 characters"`
	ConfirmNewPassword string `json:"confirm_new_password" form:"confirm_new_password" valid:"required~Confirm new password can't be empty, length(8|255)~Minimum password is 8 characters"`
}

type UserModifyPayload struct {
	Username  string    `json:"username" form:"username"`
	FullName  string    `json:"full_name" form:"full_name"`
	Phone     string    `json:"phone" form:"phone"`
	BirthDate time.Time `json:"birth_date" form:"birth_date"`
}
