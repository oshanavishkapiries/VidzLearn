package dbiface

import "context"

type Database interface {
	FindOne(ctx context.Context, collection string, filter interface{}) (interface{}, error)
	FindMany(ctx context.Context, collection string, filter interface{}) ([]interface{}, error)
	InsertOne(ctx context.Context, collection string, data interface{}) error
	UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) error
	DeleteOne(ctx context.Context, collection string, filter interface{}) error
	Aggregate(ctx context.Context, collection string, pipeline interface{}) ([]interface{}, error)
}

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string, ttlSeconds int) error
	Delete(key string) error
	Exists(key string) (bool, error)
}
