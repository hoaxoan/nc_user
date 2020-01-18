package usecase

import (
	"context"
	"log"
	"github.com/hoaxoan/nc_user/user"
	"github.com/hoaxoan/nc_user/user/service"
	"github.com/hoaxoan/nc_user/model"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	Repo         user.Repository
	TokenService service.Authable
}

func NewUserUsecase(repo user.Repository, tokenService service.Authable) user.Usecase {
	return &userUsecase{
		Repo:    repo,
		TokenService: tokenService,
	}
}

func (ucase *userUsecase) Get(ctx context.Context, req *model.User, res *model.UserResponse) error {
	user, err := ucase.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (ucase *userUsecase) GetAll(ctx context.Context, req *model.UserRequest, res *model.UserResponse) error {
	users, err := ucase.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (ucase *userUsecase) Create(ctx context.Context, req *model.User, res *model.UserResponse) error {

	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	if err := ucase.Repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (ucase *userUsecase) Update(ctx context.Context, req *model.User, res *model.UserResponse) error {
	existUser, err := ucase.Repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	if (req.Password != ""){
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		req.Password = string(hashedPass)
	} else if (existUser.Password != "") {		
		req.Password = existUser.Password
	}
	
	if err := ucase.Repo.Update(req); err != nil {
		return err
	}
	res.User = req
	return nil
}


func (ucase *userUsecase) Auth(ctx context.Context, req *model.User, res *model.Token) error {
	log.Println("Logging in with:", req.Email, req.Password)
	user, err := ucase.Repo.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return err
	}

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := ucase.TokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (ucase *userUsecase) ValidateToken(ctx context.Context, req *model.Token, res *model.Token) error {

	// Decode token
	claims, err := ucase.TokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	log.Println(claims)

	// if claims.User.Id == 0 {
	// 	return errors.New("invalid user")
	// }

	res.Valid = true

	return nil
}
