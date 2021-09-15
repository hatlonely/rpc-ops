package ops

import (
	"context"
	"fmt"
	"time"

	"github.com/hatlonely/go-kit/micro"
	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/wrap"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver"
)

func init() {
	micro.RegisterRetryRetryIf("mongo", func(err error) bool {
		switch e := err.(type) {
		case driver.Error:
			if e.NetworkError() || e.NodeIsRecovering() || e.NodeIsShuttingDown() {
				return true
			}
		}
		return false
	})

	wrap.RegisterErrCode(driver.Error{}, func(err error) string {
		return err.(driver.Error).Name
	})
	wrap.RegisterErrCode(mongo.WriteException{}, func(err error) string {
		for _, me := range err.(mongo.WriteException).WriteErrors {
			return fmt.Sprintf("E%d", me.Code)
		}
		return "Unknown"
	})
}

type ManagerOptions struct {
	Mongo                wrap.MongoClientWrapperOptions
	Database             string        `dft:"ops"`
	RepositoryCollection string        `dft:"repository"`
	VariableCollection   string        `dft:"variable"`
	SequenceCollection   string        `dft:"sequence"`
	JobCollection        string        `dft:"job"`
	Timeout              time.Duration `dft:"1s"`
	JobExpiration        time.Duration `dft:"72h"`
}

type Manager struct {
	client  *wrap.MongoClientWrapper
	options *ManagerOptions
}

func NewManagerWithOptions(options *ManagerOptions, opts ...refx.Option) (*Manager, error) {
	client, err := wrap.NewMongoClientWrapperWithOptions(&options.Mongo, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "wrap.NewMongoClientWrapperWithOptions failed")
	}

	{
		collection := client.Database(options.Database).Collection(options.RepositoryCollection)
		ctx, cancel := context.WithTimeout(context.Background(), options.Timeout)
		defer cancel()
		if _, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
			{Keys: bson.M{"endpoint": 1, "team": 1, "name": 1}, Options: mopt.Index().SetUnique(true)},
		}); err != nil {
			return nil, errors.Wrap(err, "mongo.Indexes.CreateMany failed")
		}
	}
	{
		collection := client.Database(options.Database).Collection(options.VariableCollection)
		ctx, cancel := context.WithTimeout(context.Background(), options.Timeout)
		defer cancel()
		if _, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
			{Keys: bson.M{"name": 1}, Options: mopt.Index().SetUnique(true).SetName("nameIdx")},
		}); err != nil {
			return nil, errors.Wrap(err, "mongo.Indexes.CreateMany failed")
		}
	}
	{
		collection := client.Database(options.Database).Collection(options.JobCollection)
		mongoCtx, cancel := context.WithTimeout(context.Background(), options.Timeout)
		defer cancel()
		if _, err := collection.Indexes().CreateMany(mongoCtx, []mongo.IndexModel{
			{Keys: bson.M{"jobID": 1}, Options: mopt.Index().SetName("jobIDIdx")},
			{Keys: bson.M{"createAt": 1}, Options: mopt.Index().SetName("createAtIdx").SetExpireAfterSeconds(int32(options.JobExpiration.Seconds()))},
			{Keys: bson.M{"state": 1}, Options: mopt.Index().SetName("stateIdx")},
		}); err != nil {
			return nil, errors.Wrap(err, "mongo.Indexes.CreateMany failed")
		}
	}
	{
		collection := client.Database(options.Database).Collection(options.SequenceCollection)
		ctx, cancel := context.WithTimeout(context.Background(), options.Timeout)
		defer cancel()
		if _, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
			{Keys: bson.M{"key": 1}, Options: mopt.Index().SetUnique(true).SetName("keyIdx")},
		}); err != nil {
			return nil, errors.Wrap(err, "mongo.Indexes.CreateMany failed")
		}
	}

	return &Manager{
		client:  client,
		options: options,
	}, nil
}
