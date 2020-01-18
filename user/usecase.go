package user

import (
	"context"
	"github.com/hoaxoan/nc_user/model"
)

type Usecase interface {
	Get(ctx context.Context, req *model.User, res *model.UserResponse) error
	GetAll(ctx context.Context, req *model.UserRequest, res *model.UserResponse) error
	Create(ctx context.Context, req *model.User, res *model.UserResponse) error
	Update(ctx context.Context, req *model.User, res *model.UserResponse) error
	Auth(ctx context.Context, req *model.User, res *model.Token) error
	ValidateToken(ctx context.Context, req *model.Token, res *model.Token) error
}
