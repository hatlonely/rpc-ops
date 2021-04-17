package service

import (
	"github.com/hatlonely/go-kit/refx"
	"github.com/pkg/errors"

	"github.com/hatlonely/rpc-ops/api/gen/go/api"
	"github.com/hatlonely/rpc-ops/internal/storage"
)

type Options struct {
	Storage storage.OpsOptions
}

func NewServiceWithOptions(options *Options, opts ...refx.Option) (*Service, error) {
	storage, err := storage.NewOpsStorageWithOptions(&options.Storage, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "storage.NewOpsStorageWithOptions failed")
	}

	return &Service{
		storage: storage,
		options: options,
	}, nil
}

type Service struct {
	storage *storage.OpsStorage
	options *Options

	api.UnimplementedOpsServiceServer
}
