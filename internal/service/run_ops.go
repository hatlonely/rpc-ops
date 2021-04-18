package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/hatlonely/rpc-ops/api/gen/go/api"
	"github.com/hatlonely/rpc-ops/internal/ops"
)

func (s *Service) RunOps(ctx context.Context, req *api.RunOpsReq) (*api.RunOpsRes, error) {
	id, err := s.manager.PutJob(ctx, &ops.Job{
		RepositoryID: req.RepositoryID,
		VariableID:   req.VariableID,
		Version:      req.Version,
		State:        ops.JobStateWaiting,
	})
	if err != nil {
		return nil, errors.WithMessage(err, "s.manager.PutJob failed")
	}

	return &api.RunOpsRes{JobID: id}, nil
}
