package ops

import (
	"context"
	"time"

	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/wrap"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

type Options struct {
	Mongo                wrap.MongoClientWrapperOptions
	Database             string        `dft:"ops"`
	RepositoryCollection string        `dft:"repository"`
	VariableCollection   string        `dft:"variable"`
	JobCollection        string        `dft:"job"`
	Timeout              time.Duration `dft:"1s"`
	JobExpiration        time.Duration `dft:"72h"`
}

type Manager struct {
	client  *wrap.MongoClientWrapper
	options *Options
}

func NewManagerWithOptions(options *Options, opts ...refx.Option) (*Manager, error) {
	client, err := wrap.NewMongoClientWrapperWithOptions(&options.Mongo, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "wrap.NewMongoClientWrapperWithOptions failed")
	}

	{
		collection := client.Database(options.Database).Collection(options.RepositoryCollection)
		ctx, cancel := context.WithTimeout(context.Background(), options.Timeout)
		defer cancel()
		if _, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
			{Keys: bson.M{"endpoint": 1, "name": 1}, Options: mopt.Index().SetUnique(true)},
		}); err != nil {
			return nil, errors.Wrap(err, "mongo.Indexes.CreateMany failed")
		}
	}

	return &Manager{
		client:  client,
		options: options,
	}, nil
}
