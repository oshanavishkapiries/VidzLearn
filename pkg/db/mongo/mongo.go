package mongo

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/Cenzios/pf-backend/pkg/db/dbiface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongoImpl implements dbiface.Database
// It wraps a mongo.Client and mongo.Database

type mongoImpl struct {
	client *mongo.Client
	db     *mongo.Database
}

// New returns a new mongoImpl as dbiface.Database
func New() dbiface.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).
		SetMaxPoolSize(50).
		SetMinPoolSize(5).
		SetMaxConnIdleTime(10 * time.Minute)

	client, _ := mongo.Connect(ctx, clientOpts)
	dbName := os.Getenv("MONGO_DB")
	return &mongoImpl{client: client, db: client.Database(dbName)}
}

func (m *mongoImpl) FindOne(ctx context.Context, coll string, filter interface{}) (interface{}, error) {
	var result bson.M
	err := m.db.Collection(coll).FindOne(ctx, filter).Decode(&result)
	return result, err
}

func (m *mongoImpl) FindMany(ctx context.Context, coll string, filter interface{}) ([]interface{}, error) {
	cursor, err := m.db.Collection(coll).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []interface{}
	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		results = append(results, doc)
	}
	return results, nil
}

func (m *mongoImpl) InsertOne(ctx context.Context, coll string, data interface{}) error {
	_, err := m.db.Collection(coll).InsertOne(ctx, data)
	return err
}

func (m *mongoImpl) UpdateOne(ctx context.Context, coll string, filter interface{}, update interface{}) error {
	res, err := m.db.Collection(coll).UpdateOne(ctx, filter, bson.M{"$set": update})
	if res.MatchedCount == 0 {
		return errors.New("no document matched")
	}
	return err
}

func (m *mongoImpl) DeleteOne(ctx context.Context, coll string, filter interface{}) error {
	_, err := m.db.Collection(coll).DeleteOne(ctx, filter)
	return err
}

func (m *mongoImpl) Aggregate(ctx context.Context, coll string, pipeline interface{}) ([]interface{}, error) {
	cursor, err := m.db.Collection(coll).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []interface{}
	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		results = append(results, doc)
	}
	return results, nil
}
