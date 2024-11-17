package model

import "encoding/json"

type TokenResponse struct {
	Token string `json:"token"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ToJSON(response any) []byte {
	res, _ := json.Marshal(response)
	return res
}

const (
	Header      = "web/template/components/_header.html"
	Navbar      = "web/template/components/_navbar.html"
	Footer      = "web/template/components/_footer.html"
	UserNavbar  = "web/template/components/_navbar_user.html"
	AdminNavbar = "web/template/components/_navbar_admin.html"
)
