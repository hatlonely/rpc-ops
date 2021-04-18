package ops

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

type Sequence struct {
	Key string `bson:"key"`
	Val int32  `bson:"val"`
}

func (s *Manager) IncSequence(ctx context.Context, key string) (int32, error) {
	collection := s.client.Database(s.options.Database).Collection(s.options.SequenceCollection)
	var sequence Sequence
	if err := collection.FindOneAndUpdate(
		ctx, bson.M{"key": key}, bson.M{"$inc": bson.M{"val": 1}},
		mopt.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(mopt.After),
	).Decode(&sequence); err != nil {
		return 0, errors.Wrap(err, "mongo.Collection.FindOneAndUpdate failed")
	}
	return sequence.Val, nil
}

func (s *Manager) DelSequence(ctx context.Context, key string) error {
	collection := s.client.Database(s.options.Database).Collection(s.options.SequenceCollection)
	if _, err := collection.DeleteOne(ctx, bson.M{"key": key}); err != nil {
		return errors.Wrap(err, "mongo.Collection.DeleteOne failed")
	}
	return nil
}
