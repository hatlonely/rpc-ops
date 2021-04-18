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

type Job struct {
	ID            string    `json:"id" bson:"_id,omitempty"`
	Seq           int32     `json:"seq,omitempty" bson:"seq,omitempty"`
	State         string    `json:"state,omitempty" bson:"state,omitempty"`
	RepositoryID  string    `json:"repositoryID,omitempty" bson:"repositoryID,omitempty"`
	VariableID    string    `json:"variableID,omitempty" bson:"variableID,omitempty"`
	Version       string    `json:"version,omitempty" bson:"version,omitempty"`
	CreateAt      time.Time `json:"createAt,omitempty" bson:"createAt,omitempty"`
	UpdateAt      time.Time `json:"updateAt,omitempty" bson:"updateAt,omitempty"`
	ScheduleAt    time.Time `json:"scheduleAt,omitempty" bson:"scheduleAt,omitempty"`
	ElapseSeconds int32     `json:"elapseSeconds,omitempty" bson:"elapseSeconds,omitempty"`
	Output        string    `json:"output,omitempty" bson:"output,omitempty"`
}

func (s *Manager) GetJob(ctx context.Context, id string) (*Job, error) {
	collection := s.client.Database(s.options.Database).Collection(s.options.JobCollection)
	var job Job
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&job); err != nil {
		return nil, errors.Wrapf(err, "mongo.Collection.FindOne failed. id: [%v]", id)
	}
	return &job, nil
}

func (s *Manager) DelJob(ctx context.Context, id string) error {
	collection := s.client.Database(s.options.Database).Collection(s.options.JobCollection)
	ctx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return errors.Wrapf(err, "mongo.Collection.DeleteOne failed. id: [%v]", id)
	}
	return nil
}

func (s *Manager) PutJob(ctx context.Context, job *Job) (string, error) {
	seq, err := s.IncSequence(ctx, job.RepositoryID)
	if err != nil {
		return "", errors.Wrap(err, "s.IncSequence failed")
	}
	collection := s.client.Database(s.options.Database).Collection(s.options.JobCollection)
	ctx, cancel := context.WithTimeout(ctx, s.options.Timeout)
	defer cancel()
	buf := make([]byte, 32)
	hex.Encode(buf, uuid.NewV4().Bytes())
	job.ID = string(buf)
	job.CreateAt = time.Now()
	job.UpdateAt = time.Now()
	job.Seq = seq
	res, err := collection.InsertOne(ctx, job)
	if err != nil {
		return "", errors.Wrap(err, "mongo.Collection.InsertOne failed")
	}
	return res.InsertedID.(string), nil
}

func (s *Manager) UpdateJob(ctx context.Context, job *Job) error {
	collection := s.client.Database(s.options.Database).Collection(s.options.JobCollection)
	job.UpdateAt = time.Now()
	_, err := collection.UpdateOne(ctx, bson.M{"_id": job.ID}, bson.M{"$set": job})
	if err != nil {
		return errors.Wrapf(err, "mongo.Collection.UpdateOne failed. id: [%v]", job.ID)
	}
	return nil
}

func (s *Manager) ListJob(ctx context.Context, offset int32, limit int32) ([]*Job, error) {
	collection := s.client.Database(s.options.Database).Collection(s.options.JobCollection)
	res, err := collection.Find(ctx, bson.M{}, mopt.Find().SetLimit(int64(limit)).SetSkip(int64(offset)).SetSort(bson.M{"createAt": -1}))
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find failed")
	}
	defer res.Close(ctx)
	var repositories []*Job
	if err := res.All(ctx, &repositories); err != nil {
		return nil, errors.Wrap(err, "mongo.Collection.Find.All failed")
	}
	return repositories, nil
}
