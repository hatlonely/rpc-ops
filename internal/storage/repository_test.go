package storage

import (
	"context"
	"testing"
	"time"

	"github.com/hatlonely/go-kit/micro"
	"github.com/hatlonely/go-kit/wrap"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
)

func TestRepository(t *testing.T) {
	Convey("TestRepository", t, func() {
		s, err := NewOpsStorageWithOptions(&OpsOptions{
			Database:             "ops",
			RepositoryCollection: "repository",
			Timeout:              10 * time.Second,
			Mongo: wrap.MongoClientWrapperOptions{
				Retry: micro.RetryOptions{
					Attempts: 1,
				},
				Mongo: wrap.MongoOptions{
					URI:            "mongodb://localhost:27017",
					ConnectTimeout: 3 * time.Second,
					PingTimeout:    2 * time.Second,
				},
			},
		})
		So(err, ShouldBeNil)

		_, _ = s.client.Database("ops").Collection("repository").DeleteMany(context.Background(), bson.M{
			"endpoint": "github.com",
			"name":     "rpc-ops",
		})
		_, _ = s.client.Database("ops").Collection("repository").DeleteMany(context.Background(), bson.M{
			"endpoint": "test-endpoint",
			"name":     "test-name",
		})

		var id string
		{
			id, err = s.PutRepository(context.Background(), &Repository{
				Username: "hatlonely",
				Password: "123456",
				Endpoint: "github.com",
				Name:     "rpc-ops",
			})
			So(err, ShouldBeNil)
			So(id, ShouldNotBeEmpty)

			repository, err := s.GetRepository(context.Background(), id)
			So(err, ShouldBeNil)
			So(repository.ID, ShouldEqual, id)
			So(repository.Username, ShouldEqual, "hatlonely")
			So(repository.Password, ShouldEqual, "123456")
			So(repository.Endpoint, ShouldEqual, "github.com")
			So(repository.Name, ShouldEqual, "rpc-ops")
		}

		{
			err = s.UpdateRepository(context.Background(), &Repository{
				ID:       id,
				Password: "121231",
				Endpoint: "test-endpoint",
				Name:     "test-name",
			})
			So(err, ShouldBeNil)

			repository, err := s.GetRepository(context.Background(), id)
			So(err, ShouldBeNil)
			So(repository.ID, ShouldEqual, id)
			So(repository.Username, ShouldEqual, "hatlonely")
			So(repository.Password, ShouldEqual, "121231")
			So(repository.Endpoint, ShouldEqual, "test-endpoint")
			So(repository.Name, ShouldEqual, "test-name")
		}
	})
}
