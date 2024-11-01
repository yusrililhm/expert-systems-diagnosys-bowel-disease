package repo

import (
	"usus-sehat/server/entity"
	"usus-sehat/server/exception"
	"usus-sehat/server/model"
)

type AdminRepo interface {
	FetchByUsername(username string) (*entity.User, exception.Exception)
	FetchAllUser() ([]*model.User, exception.Exception)
}
