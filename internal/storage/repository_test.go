package storage

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hatlonely/go-kit/micro"
	"github.com/hatlonely/go-kit/wrap"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

		for i := 0; i < 10; i++ {
			_, _ = s.client.Database("ops").Collection("repository").DeleteMany(context.Background(), bson.M{
				"endpoint": fmt.Sprintf("test-endpoint-%v", i),
				"name":     fmt.Sprintf("test-name-%v", i),
			})
		}

		var id string
		{
			id, err = s.PutRepository(context.Background(), &Repository{
				Username: "test-username-1",
				Password: "test-password-1",
				Endpoint: "test-endpoint-1",
				Name:     "test-name-1",
			})
			So(err, ShouldBeNil)
			So(id, ShouldNotBeEmpty)

			repository, err := s.GetRepository(context.Background(), id)
			So(err, ShouldBeNil)
			So(repository.ID, ShouldEqual, id)
			So(repository.Username, ShouldEqual, "test-username-1")
			So(repository.Password, ShouldEqual, "test-password-1")
			So(repository.Endpoint, ShouldEqual, "test-endpoint-1")
			So(repository.Name, ShouldEqual, "test-name-1")
		}

		{
			err = s.UpdateRepository(context.Background(), &Repository{
				ID:       id,
				Password: "test-password-2",
				Endpoint: "test-endpoint-2",
				Name:     "test-name-2",
			})
			So(err, ShouldBeNil)

			repository, err := s.GetRepository(context.Background(), id)
			So(err, ShouldBeNil)
			So(repository.ID, ShouldEqual, id)
			So(repository.Username, ShouldEqual, "test-username-1")
			So(repository.Password, ShouldEqual, "test-password-2")
			So(repository.Endpoint, ShouldEqual, "test-endpoint-2")
			So(repository.Name, ShouldEqual, "test-name-2")
		}

		{
			err = s.DelRepository(context.Background(), id)
			So(err, ShouldBeNil)

			repository, err := s.GetRepository(context.Background(), id)
			So(errors.Cause(err), ShouldEqual, mongo.ErrNoDocuments)
			So(repository, ShouldBeNil)
		}

		{
			var ids []string
			for i := 0; i < 10; i++ {
				id, err := s.PutRepository(context.Background(), &Repository{
					Username: fmt.Sprintf("test-username-%v", i),
					Password: fmt.Sprintf("test-password-%v", i),
					Endpoint: fmt.Sprintf("test-endpoint-%v", i),
					Name:     fmt.Sprintf("test-name-%v", i),
				})
				So(err, ShouldBeNil)
				ids = append(ids, id)
			}

			repositories, err := s.ListRepository(context.Background(), 3, 4)
			So(err, ShouldBeNil)
			for i := 0; i < 4; i++ {
				fmt.Println(repositories[i])
			}
		}
	})
}
