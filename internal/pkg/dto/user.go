package dto

type AddUserPayload struct {
	Username string
	Email    string
	FullName string
	Gender   bool
	Password string
}

type UserLoginPayload struct {
	Username string
	Password string
}

type ChangePasswordPayload struct {
	OldPassword string
	NewPassword string
}

type EditUserPayload struct {
	Username string
	Email    string
	FullName string
	Gender   bool
}
