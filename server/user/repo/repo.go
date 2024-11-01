package repo

import (
	"usus-sehat/server/entity"
	"usus-sehat/server/exception"
)

type UserRepo interface {
	Add(user *entity.User) exception.Exception
	FetchByPhone(phone string) (*entity.User, exception.Exception)
	FetchById(id int) (*entity.User, exception.Exception)
	Modify(user *entity.User) exception.Exception
	ChangePassword(user *entity.User) exception.Exception
}
