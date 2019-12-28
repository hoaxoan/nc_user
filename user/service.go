package user

import (
	"context"
	"log"

	md "github.com/hoaxoan/nc_course/nc_user/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo         *UserRepository
	TokenService Authable
}

func (srv *UserService) Get(ctx context.Context, req *md.User, res *md.UserResponse) error {
	user, err := srv.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *UserService) GetAll(ctx context.Context, req *md.UserRequest, res *md.UserResponse) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *UserService) Auth(ctx context.Context, req *md.User, res *md.Token) error {
	log.Println("Logging in with:", req.Email, req.Password)
	user, err := srv.Repo.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return err
	}

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.TokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (srv *UserService) Create(ctx context.Context, req *md.User, res *md.UserResponse) error {

	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	if err := srv.Repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (srv *UserService) Update(ctx context.Context, req *md.User, res *md.UserResponse) error {
	if err := srv.Repo.Update(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

// func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {

// 	// Decode token
// 	claims, err := srv.tokenService.Decode(req.Token)
// 	if err != nil {
// 		return err
// 	}

// 	log.Println(claims)

// 	if claims.User.Id == "" {
// 		return errors.New("invalid user")
// 	}

// 	res.Valid = true

// 	return nil
// }
