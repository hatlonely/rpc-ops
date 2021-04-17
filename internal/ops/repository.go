package ops

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Endpoint string `json:"endpoint,omitempty" bson:"endpoint,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	OpsFile  string `json:"opsFile,omitempty" bson:"opsFile,omitempty"`
	CreateAt int32  `json:"createAt,omitempty" bson:"createAt,omitempty"`
	UpdateAt int32  `json:"updateAt,omitempty" bson:"updateAt,omitempty"`
}

func (s *Manager) GetRepository(ctx context.Context, id string) (*Repository, error) {
	collection := s.client.Database(s.options.Database).Collection(s.options.RepositoryCollection)
	var repository Repository
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&repository); err != nil {
		return nil, errors.Wrapf(err, "mongo.Collection.FindOne failed. id: [%v]", id)
	}
	return &repository, nil
}

func (s *Manager) DelRepository(ctx context.Context, id string) error {
	collection := s.client.Database(s.options.Database).Collection(s.options.RepositoryCollection)
	ctx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return errors.Wrapf(err, "mongo.Collection.DeleteOne failed. id: [%v]", id)
	}
	return nil
}

func (s *Manager) PutRepository(ctx context.Context, repository *Repository) (string, error) {
	collection := s.client.Database(s.options.Database).Collection(s.options.RepositoryCollection)
	ctx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	buf := make([]byte, 32)
	hex.Encode(buf, uuid.NewV4().Bytes())
	repository.ID = string(buf)
	repository.CreateAt = int32(time.Now().Unix())
	repository.UpdateAt = int32(time.Now().Unix())
	res, err := collection.InsertOne(ctx, repository)
	if err != nil {
		return "", errors.Wrap(err, "mongo.Collection.InsertOne failed")
	}
	return res.InsertedID.(string), nil
}

func (s *Manager) UpdateRepository(ctx context.Context, repository *Repository) error {
	collection := s.client.Database(s.options.Database).Collection(s.options.RepositoryCollection)
	repository.UpdateAt = int32(time.Now().Unix())
	_, err := collection.UpdateOne(ctx, bson.M{"_id": repository.ID}, bson.M{"$set": repository})
	if err != nil {
		return errors.Wrapf(err, "mongo.Collection.UpdateOne failed. id: [%v]", repository.ID)
	}
	return nil
}

func (s *Manager) ListRepository(ctx context.Context, offset int32, limit int32) ([]*Repository, error) {
	collection := s.client.Database(s.options.Database).Collection(s.options.RepositoryCollection)
	res, err := collection.Find(ctx, bson.M{}, mopt.Find().SetLimit(int64(limit)).SetSkip(int64(offset)).SetSort(bson.M{"createAt": -1}))
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
