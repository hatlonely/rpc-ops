package ops

import (
	"context"
	"testing"
	"time"

	"github.com/hatlonely/go-kit/micro"
	"github.com/hatlonely/go-kit/wrap"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
)

func TestSequence(t *testing.T) {
	Convey("TestSequence", t, func() {
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

		_, err = s.client.Database("ops").Collection("sequence").DeleteMany(context.Background(), bson.M{})
		So(err, ShouldBeNil)

		{
			i, err := s.IncSequence(context.Background(), "key-1")
			So(err, ShouldBeNil)
			So(i, ShouldEqual, 1)
			i, err = s.IncSequence(context.Background(), "key-1")
			So(err, ShouldBeNil)
			So(i, ShouldEqual, 2)
			i, err = s.IncSequence(context.Background(), "key-1")
			So(err, ShouldBeNil)
			So(i, ShouldEqual, 3)
		}
		{
			So(s.DelSequence(context.Background(), "key-1"), ShouldBeNil)
			So(s.DelSequence(context.Background(), "key-1"), ShouldBeNil)
		}
	})
}
