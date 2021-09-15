package service

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/rpcx"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/rpc-ops/api/gen/go/api"
	"github.com/hatlonely/rpc-ops/internal/ops"
)

type Options struct {
	Ops      ops.ManagerOptions
	WorkRoot string
}

func NewServiceWithOptions(options *Options, opts ...refx.Option) (*Service, error) {
	manager, err := ops.NewManagerWithOptions(&options.Ops, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "manager.NewManagerWithOptions failed")
	}

	return &Service{
		manager: manager,
		options: options,
	}, nil
}

type Service struct {
	manager *ops.Manager
	options *Options

	api.UnimplementedOpsServiceServer
}

func (s *Service) Ping(ctx context.Context, req *api.Empty) (*api.Empty, error) {
	return &api.Empty{}, nil
}

func toRpcError(err error) error {
	fmt.Println(reflect.TypeOf(errors.Cause(err)))
	switch e := errors.Cause(err).(type) {
	case driver.Error:
		if e.NetworkError() {
			return err
		}
		return rpcx.NewError(err, codes.InvalidArgument, fmt.Sprintf("mongo.%s", e.Name), e.Message)
	case mongo.WriteException:
		for _, me := range e.WriteErrors {
			return rpcx.NewError(err, codes.InvalidArgument, fmt.Sprintf("mongo.E%d", me.Code), me.Message)
		}
	}

	return err
}
