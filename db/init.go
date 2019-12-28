package db

import (
	"context"
	"log"
	"time"

	"github.com/hoaxoan/nc_course/nc_user/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connection() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config.Mongo.Uri))
	if err != nil {
		log.Fatalf("connect error :%v", err)
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	return client, nil
}
