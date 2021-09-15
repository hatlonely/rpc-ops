package service

import (
	"context"
	"time"

	"github.com/hatlonely/rpc-ops/api/gen/go/api"
	"github.com/hatlonely/rpc-ops/internal/ops"
)

func (s *Service) ListRepository(ctx context.Context, req *api.ListRepositoryReq) (*api.ListRepositoryRes, error) {
	repos, err := s.manager.ListRepository(ctx, req.Offset, req.Limit)
	if err != nil {
		return nil, toRpcError(err)
	}
	var res api.ListRepositoryRes
	for _, repo := range repos {
		res.Repositories = append(res.Repositories, &api.Repository{
			Id:       repo.ID,
			Username: repo.Username,
			Password: repo.Password,
			Endpoint: repo.Endpoint,
			Team:     repo.Team,
			Name:     repo.Name,
			Playbook: repo.Playbook,
			CreateAt: repo.CreateAt.Format(time.RFC3339),
			UpdateAt: repo.UpdateAt.Format(time.RFC3339),
		})
	}
	return &res, nil
}

func (s *Service) GetRepository(ctx context.Context, req *api.RepositoryID) (*api.Repository, error) {
	repo, err := s.manager.GetRepository(ctx, req.Id)
	if err != nil {
		return nil, toRpcError(err)
	}
	return &api.Repository{
		Id:       repo.ID,
		Username: repo.Username,
		Password: repo.Password,
		Endpoint: repo.Endpoint,
		Team:     repo.Team,
		Name:     repo.Name,
		Playbook: repo.Playbook,
		CreateAt: repo.CreateAt.Format(time.RFC3339),
		UpdateAt: repo.UpdateAt.Format(time.RFC3339),
	}, nil
}

func (s *Service) DelRepository(ctx context.Context, req *api.RepositoryID) (*api.Empty, error) {
	err := s.manager.DelRepository(ctx, req.Id)
	if err != nil {
		return nil, toRpcError(err)
	}
	return &api.Empty{}, nil
}

func (s *Service) PutRepository(ctx context.Context, req *api.Repository) (*api.RepositoryID, error) {
	id, err := s.manager.PutRepository(ctx, &ops.Repository{
		Username: req.Username,
		Password: req.Password,
		Endpoint: req.Endpoint,
		Team:     req.Team,
		Name:     req.Name,
		Playbook: req.Playbook,
	})
	if err != nil {
		return nil, toRpcError(err)
	}
	return &api.RepositoryID{Id: id}, nil
}

func (s *Service) UpdateRepository(ctx context.Context, req *api.Repository) (*api.Empty, error) {
	err := s.manager.UpdateRepository(ctx, &ops.Repository{
		ID:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Endpoint: req.Endpoint,
		Team:     req.Team,
		Name:     req.Name,
		Playbook: req.Playbook,
	})
	if err != nil {
		return nil, toRpcError(err)
	}
	return &api.Empty{}, nil
}
