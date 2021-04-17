package storage

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hatlonely/go-kit/micro"
	"github.com/hatlonely/go-kit/wrap"
	. "github.com/smartystreets/goconvey/convey"
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

		id, err := s.PutRepository(context.Background(), &Repository{
			Username: "hatlonely",
			Password: "123456",
			Endpoint: "github.com",
			Name:     "rpc-ops",
		})
		So(err, ShouldBeNil)
		fmt.Println(id)
	})
}
