package user

import (
	"github.com/hoaxoan/nc_user/model"
)


type Repository interface {
	GetAll() ([]*model.User, error)
	Get(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
}