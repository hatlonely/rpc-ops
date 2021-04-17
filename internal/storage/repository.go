package storage

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

var ErrInvalidObjectID = errors.New("invalid object id")

type Repository struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Endpoint string `json:"endpoint"`
	Name     string `json:"repository"`
	CreateAt int32  `json:"createAt"`
	UpdateAt int32  `json:"updateAt"`
}

func (s *OpsStorage) GetRepository(ctx context.Context, id string) (*Repository, error) {
	collection := s.client.Database(s.options.Database).Collection(s.options.RepositoryCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectID
	}
	ctx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	var repository Repository
	if err := collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&repository); err != nil {
		return nil, errors.Wrapf(err, "mongo.Collection.FindOne failed. id: [%v]", id)
	}
	return &repository, nil
}

func (s *OpsStorage) DelRepository(ctx context.Context, id string) error {
	collection := s.client.Database(s.options.Database).Collection(s.options.RepositoryCollection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidObjectID
	}
	ctx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return errors.Wrapf(err, "mongo.Collection.DeleteOne failed. id: [%v]", id)
	}
	return nil
}

func (s *OpsStorage) PutRepository(ctx context.Context, repository *Repository) (string, error) {
	collection := s.client.Database(s.options.Database).Collection(s.options.RepositoryCollection)
	ctx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	repository.CreateAt = int32(time.Now().Unix())
	repository.UpdateAt = int32(time.Now().Unix())
	res, err := collection.InsertOne(ctx, repository)
	if err != nil {
		return "", errors.Wrap(err, "mongo.Collection.InsertOne failed")
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (s *OpsStorage) UpdateRepository(ctx context.Context, repository *Repository) error {
	collection := s.client.Database(s.options.Database).Collection(s.options.RepositoryCollection)
	objectID, err := primitive.ObjectIDFromHex(repository.ID)
	if err != nil {
		return ErrInvalidObjectID
	}
	ctx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	repository.UpdateAt = int32(time.Now().Unix())
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": repository})
	if err != nil {
		return errors.Wrapf(err, "mongo.Collection.UpdateOne failed. id: [%v]", repository.ID)
	}
	return nil
}

func (s *OpsStorage) ListRepository(ctx context.Context, offset int32, limit int32) ([]*Repository, error) {
	collection := s.client.Database(s.options.Database).Collection(s.options.RepositoryCollection)
	ctx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	res, err := collection.Find(ctx, bson.M{}, mopt.Find().SetLimit(int64(limit)).SetSkip(int64(offset)))
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find failed")
	}
	defer res.Close(ctx)
	var repositories []*Repository
	if err := res.All(ctx, &repositories); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find.All failed")
	}
	return repositories, nil
}
