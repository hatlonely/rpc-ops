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

func TestJob(t *testing.T) {
	Convey("TestJob", t, func() {
		s, err := NewManagerWithOptions(&ManagerOptions{
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

		_, err = s.client.Database("ops").Collection("job").DeleteMany(context.Background(), bson.M{})
		So(err, ShouldBeNil)

		var id string
		{
			id, err = s.PutJob(context.Background(), &Job{
				State:         JobStateWaiting,
				RepositoryID:  "test-repository-id-1",
				VariableID:    "test-variable-id-1",
				Version:       "test-version-1",
				ElapseSeconds: 2,
				Output:        "test-output-1",
			})
			So(err, ShouldBeNil)
			So(id, ShouldNotBeEmpty)

			job, err := s.GetJob(context.Background(), id)
			So(err, ShouldBeNil)
			So(job.ID, ShouldEqual, id)
			So(job.State, ShouldEqual, "Waiting")
			So(job.RepositoryID, ShouldEqual, "test-repository-id-1")
			So(job.VariableID, ShouldEqual, "test-variable-id-1")
			So(job.Version, ShouldEqual, "test-version-1")
			So(job.ElapseSeconds, ShouldEqual, 2)
			So(job.Output, ShouldEqual, "test-output-1")
		}

		{
			err = s.UpdateJob(context.Background(), &Job{
				ID:            id,
				Version:       "test-version-2",
				ElapseSeconds: 3,
				Output:        "test-output-2",
			})
			So(err, ShouldBeNil)

			job, err := s.GetJob(context.Background(), id)
			So(err, ShouldBeNil)
			So(job.ID, ShouldEqual, id)
			So(job.State, ShouldEqual, "Waiting")
			So(job.RepositoryID, ShouldEqual, "test-repository-id-1")
			So(job.VariableID, ShouldEqual, "test-variable-id-1")
			So(job.Version, ShouldEqual, "test-version-2")
			So(job.ElapseSeconds, ShouldEqual, 3)
			So(job.Output, ShouldEqual, "test-output-2")
		}

		{
			job, err := s.FindOneWaitingJob(context.Background())
			So(err, ShouldBeNil)
			So(job.ID, ShouldEqual, id)
			So(job.State, ShouldEqual, "Waiting")
			So(job.RepositoryID, ShouldEqual, "test-repository-id-1")
			So(job.VariableID, ShouldEqual, "test-variable-id-1")
			So(job.Version, ShouldEqual, "test-version-2")
			So(job.ElapseSeconds, ShouldEqual, 3)
			So(job.Output, ShouldEqual, "test-output-2")
		}

		{
			err = s.DelJob(context.Background(), id)
			So(err, ShouldBeNil)

			job, err := s.GetJob(context.Background(), id)
			So(errors.Cause(err), ShouldEqual, mongo.ErrNoDocuments)
			So(job, ShouldBeNil)
		}

		{
			var ids []string
			for i := 0; i < 10; i++ {
				id, err := s.PutJob(context.Background(), &Job{
					RepositoryID: fmt.Sprintf("test-repository-%v", i),
					VariableID:   fmt.Sprintf("test-variable-%v", i),
				})
				So(err, ShouldBeNil)
				ids = append(ids, id)
			}

			repositories, err := s.ListJob(context.Background(), 3, 4)
			So(err, ShouldBeNil)
			for i := 0; i < 4; i++ {
				fmt.Println(repositories[i])
			}
		}
	})
}
