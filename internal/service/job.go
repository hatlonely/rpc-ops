package service

import (
	"context"
	"time"

	"github.com/hatlonely/rpc-ops/api/gen/go/api"
)

func (s *Service) ListJob(ctx context.Context, req *api.ListJobReq) (*api.ListJobRes, error) {
	jobs, err := s.manager.ListJob(ctx, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}
	var res api.ListJobRes
	for _, job := range jobs {
		res.Jobs = append(res.Jobs, &api.Job{
			Id:            job.ID,
			Seq:           job.Seq,
			State:         job.State,
			RepositoryID:  job.RepositoryID,
			VariableID:    job.VariableID,
			Version:       job.Version,
			CreateAt:      job.CreateAt.Format(time.RFC3339),
			UpdateAt:      job.UpdateAt.Format(time.RFC3339),
			ScheduleAt:    job.ScheduleAt.Format(time.RFC3339),
			ElapseSeconds: job.ElapseSeconds,
			Output:        job.Output,
		})
	}
	return &res, nil
}

func (s *Service) GetJob(ctx context.Context, req *api.JobID) (*api.Job, error) {
	job, err := s.manager.GetJob(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.Job{
		Id:            job.ID,
		Seq:           job.Seq,
		State:         job.State,
		RepositoryID:  job.RepositoryID,
		VariableID:    job.VariableID,
		Version:       job.Version,
		CreateAt:      job.CreateAt.Format(time.RFC3339),
		UpdateAt:      job.UpdateAt.Format(time.RFC3339),
		ScheduleAt:    job.ScheduleAt.Format(time.RFC3339),
		ElapseSeconds: job.ElapseSeconds,
		Output:        job.Output,
	}, nil
}

func (s *Service) DelJob(ctx context.Context, req *api.JobID) (*api.Empty, error) {
	err := s.manager.DelJob(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.Empty{}, nil
}
