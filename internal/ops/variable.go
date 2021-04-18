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

type Variable struct {
	ID       string    `json:"id" bson:"_id,omitempty"`
	Name     string    `json:"name,omitempty" bson:"name,omitempty"`
	Body     string    `json:"body,omitempty" bson:"body,omitempty"`
	CreateAt time.Time `json:"createAt,omitempty" bson:"createAt,omitempty"`
	UpdateAt time.Time `json:"updateAt,omitempty" bson:"updateAt,omitempty"`
}

func (s *Manager) GetVariable(ctx context.Context, id string) (*Variable, error) {
	collection := s.client.Database(s.options.Database).Collection(s.options.VariableCollection)
	var variable Variable
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&variable); err != nil {
		return nil, errors.Wrapf(err, "mongo.Collection.FindOne failed. id: [%v]", id)
	}
	return &variable, nil
}

func (s *Manager) DelVariable(ctx context.Context, id string) error {
	collection := s.client.Database(s.options.Database).Collection(s.options.VariableCollection)
	ctx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return errors.Wrapf(err, "mongo.Collection.DeleteOne failed. id: [%v]", id)
	}
	return nil
}

func (s *Manager) PutVariable(ctx context.Context, variable *Variable) (string, error) {
	collection := s.client.Database(s.options.Database).Collection(s.options.VariableCollection)
	ctx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	buf := make([]byte, 32)
	hex.Encode(buf, uuid.NewV4().Bytes())
	variable.ID = string(buf)
	variable.CreateAt = time.Now()
	variable.UpdateAt = time.Now()
	res, err := collection.InsertOne(ctx, variable)
	if err != nil {
		return "", errors.Wrap(err, "mongo.Collection.InsertOne failed")
	}
	return res.InsertedID.(string), nil
}

func (s *Manager) UpdateVariable(ctx context.Context, variable *Variable) error {
	collection := s.client.Database(s.options.Database).Collection(s.options.VariableCollection)
	variable.UpdateAt = time.Now()
	_, err := collection.UpdateOne(ctx, bson.M{"_id": variable.ID}, bson.M{"$set": variable})
	if err != nil {
		return errors.Wrapf(err, "mongo.Collection.UpdateOne failed. id: [%v]", variable.ID)
	}
	return nil
}

func (s *Manager) ListVariable(ctx context.Context, offset int32, limit int32) ([]*Variable, error) {
	collection := s.client.Database(s.options.Database).Collection(s.options.VariableCollection)
	res, err := collection.Find(ctx, bson.M{}, mopt.Find().SetLimit(int64(limit)).SetSkip(int64(offset)).SetSort(bson.M{"createAt": -1}))
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find failed")
	}
	defer res.Close(ctx)
	var variables []*Variable
	if err := res.All(ctx, &variables); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find.All failed")
	}
	return variables, nil
}
