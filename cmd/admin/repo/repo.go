package repo

import (
	"usus-sehat/internal/domain/entity"
	"usus-sehat/internal/domain/model"
	"usus-sehat/internal/pkg/exception"
)

type AdminRepo interface {
	FetchByUsername(username string) (*entity.User, exception.Exception)
	FetchAllUser() ([]*model.User, exception.Exception)
}
