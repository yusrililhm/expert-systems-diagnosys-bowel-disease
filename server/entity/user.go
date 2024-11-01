package entity

import (
	"os"
	"strings"
	"time"
	"usus-sehat/server/exception"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `gorm:"not null; index:idx_username, unique"`
	FullName  string    `gorm:"not null"`
	Phone     string    `gorm:"not null; size:15; index:idx_phone, unique"`
	Gender    bool      `gorm:"not null"`
	BirthDate time.Time `gorm:"not null"`
	Role      string    `gorm:"not null;size:5"`
	Password  string    `gorm:"not null"`
}

func (u *User) GenerateFromPassword() {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashPassword)
}

func (u *User) CompareHashPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) ValidateToken(bearerToken string) exception.Exception {

	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	if !isBearer {
		return exception.NewUnauthenticatedError("Invalid token")
	}

	tokenFields := strings.Fields(bearerToken)

	if len(tokenFields) != 2 {
		return exception.NewUnauthenticatedError("Invalid token")
	}

	tokenString := tokenFields[1]

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exception.NewUnauthenticatedError("Invalid token")
		}

		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return exception.NewUnauthenticatedError("Invalid token")
	}

	mapClaims := jwt.MapClaims{}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return exception.NewUnauthenticatedError("Invalid token")
	} else {
		mapClaims = claims
	}

	id, ok := mapClaims["id"].(float64)

	if !ok {
		return exception.NewUnauthenticatedError("Invalid token")
	}

	u.ID = uint(id)

	username, ok := mapClaims["username"].(string)

	if !ok {
		return exception.NewUnauthenticatedError("Invalid token")
	}

	u.Username = username

	phone, ok := mapClaims["phone"].(string)

	if !ok {
		return exception.NewUnauthenticatedError("Invalid token")
	}

	u.Phone = phone

	role, ok := mapClaims["role"].(string)

	if !ok {
		return exception.NewUnauthenticatedError("Invalid token")
	}

	u.Role = role

	return nil
}

func (u *User) GenerateTokenString() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       u.ID,
		"username": u.Username,
		"phone":    u.Phone,
		"role":     u.Role,
	})

	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	return tokenString
}
