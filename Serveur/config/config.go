package config

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var userCollection *mongo.Collection

var accessDuration = time.Minute
var refreshDuration = time.Hour * 24
var sessionExpiredDuration = time.Hour * 24 * 7

func SetClient(c *mongo.Client) {
	client = c
}

func GetClient() *mongo.Client {
	return client
}

func SetUserCollection(c *mongo.Collection) {
	userCollection = c
}

func GetUserCollection() *mongo.Collection {
	return userCollection
}

func GetAccessDuration() time.Duration {
	return accessDuration
}

func GetRefreshDuration() time.Duration {
	return refreshDuration
}

func GetSessionExpiredDuration() time.Duration {
	return sessionExpiredDuration
}
