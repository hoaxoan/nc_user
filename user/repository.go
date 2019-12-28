package user

import (
	"context"

	md "github.com/hoaxoan/nc_course/nc_user/model"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	DbName  = "nc_student"
	ColName = "user"
)

type Repository interface {
	GetAll() ([]*md.User, error)
	Get(id int) (*md.User, error)
	GetByEmail(email string) (md.User, error)
	Create(user *md.User) error
	Update(user *md.User) error
}

type UserRepository struct {
	Client *mongo.Client
}

func (repo *UserRepository) collection() *mongo.Collection {
	return repo.Client.Database(DbName).Collection(ColName)
}

func (repo *UserRepository) GetAll() ([]*md.User, error) {
	var users []*md.User
	cur, err := repo.collection().Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cur.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(id int) (*md.User, error) {
	var user *md.User
	user.Id = id
	filter := bson.M{"id": id}
	err := repo.collection().FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*md.User, error) {
	var user *md.User
	filter := bson.M{"email": email}
	err := repo.collection().FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *md.User) error {
	_, err := repo.collection().InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Update(user *md.User) error {
	filter := bson.M{"email": user.Email}
	_, err := repo.collection().UpdateOne(context.TODO(), filter, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}
