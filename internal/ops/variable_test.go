package ops

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

func TestVariable(t *testing.T) {
	Convey("TestVariable", t, func() {
		s, err := NewManagerWithOptions(&Options{
			Database:             "ops",
			RepositoryCollection: "repository",
			VariableCollection:   "variable",
			SequenceCollection:   "sequence",
			JobCollection:        "job",
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
				"name": fmt.Sprintf("test-name-%v", i),
			})
		}

		var id string
		{
			id, err = s.PutVariable(context.Background(), &Variable{
				Name: "test-name-1",
				Body: "test-body-1",
			})
			So(err, ShouldBeNil)
			So(id, ShouldNotBeEmpty)

			repository, err := s.GetVariable(context.Background(), id)
			So(err, ShouldBeNil)
			So(repository.ID, ShouldEqual, id)
			So(repository.Name, ShouldEqual, "test-name-1")
			So(repository.Body, ShouldEqual, "test-body-1")
		}

		{
			err = s.UpdateVariable(context.Background(), &Variable{
				ID:   id,
				Body: "test-body-2",
			})
			So(err, ShouldBeNil)

			repository, err := s.GetVariable(context.Background(), id)
			So(err, ShouldBeNil)
			So(repository.ID, ShouldEqual, id)
			So(repository.Name, ShouldEqual, "test-name-1")
			So(repository.Body, ShouldEqual, "test-body-2")
		}

		{
			err = s.DelVariable(context.Background(), id)
			So(err, ShouldBeNil)

			repository, err := s.GetVariable(context.Background(), id)
			So(errors.Cause(err), ShouldEqual, mongo.ErrNoDocuments)
			So(repository, ShouldBeNil)
		}

		{
			var ids []string
			for i := 0; i < 10; i++ {
				id, err := s.PutVariable(context.Background(), &Variable{
					Name: fmt.Sprintf("test-name-%v", i),
					Body: fmt.Sprintf("test-body-%v", i),
				})
				So(err, ShouldBeNil)
				ids = append(ids, id)
			}

			repositories, err := s.ListVariable(context.Background(), 3, 4)
			So(err, ShouldBeNil)
			for i := 0; i < 4; i++ {
				fmt.Println(repositories[i])
			}
		}
	})
}
