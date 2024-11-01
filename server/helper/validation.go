package helper

import (
	"usus-sehat/server/exception"

	"github.com/asaskevich/govalidator"
)

func ValidatePayload(payload any) exception.Exception {
	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return nil
	}
	
	return nil
}
