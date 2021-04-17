package service

import (
	"github.com/hatlonely/go-kit/refx"
	"github.com/pkg/errors"

	"github.com/hatlonely/rpc-ops/api/gen/go/api"
	"github.com/hatlonely/rpc-ops/internal/ops"
)

type Options struct {
	Ops ops.Options
}

func NewServiceWithOptions(options *Options, opts ...refx.Option) (*Service, error) {
	manager, err := ops.NewOpsStorageWithOptions(&options.Ops, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "manager.NewOpsStorageWithOptions failed")
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
