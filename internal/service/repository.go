package service

import (
	"context"
	"time"

	"github.com/hatlonely/rpc-ops/api/gen/go/api"
	"github.com/hatlonely/rpc-ops/internal/storage"
)

func (s *Service) ListRepository(ctx context.Context, req *api.ListRepositoryReq) (*api.ListRepositoryRes, error) {
	repos, err := s.storage.ListRepository(ctx, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}
	var res api.ListRepositoryRes
	for _, repo := range repos {
		res.Repositories = append(res.Repositories, &api.Repository{
			Id:       repo.ID,
			Username: repo.Username,
			Password: repo.Password,
			Endpoint: repo.Endpoint,
			Name:     repo.Name,
			CreateAt: time.Unix(int64(repo.CreateAt), 0).Format(time.RFC3339),
			UpdateAt: time.Unix(int64(repo.UpdateAt), 0).Format(time.RFC3339),
		})
	}
	return &res, nil
}

func (s *Service) GetRepository(ctx context.Context, req *api.GetRepositoryReq) (*api.Repository, error) {
	repo, err := s.storage.GetRepository(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.Repository{
		Id:       repo.ID,
		Username: repo.Username,
		Password: repo.Password,
		Endpoint: repo.Endpoint,
		Name:     repo.Name,
	}, nil
}

func (s *Service) DelRepository(ctx context.Context, req *api.DelRepositoryReq) (*api.Empty, error) {
	err := s.storage.DelRepository(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.Empty{}, nil
}

func (s *Service) PutRepository(ctx context.Context, req *api.Repository) (*api.PutRepositoryRes, error) {
	id, err := s.storage.PutRepository(ctx, &storage.Repository{
		Username: req.Username,
		Password: req.Password,
		Endpoint: req.Endpoint,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &api.PutRepositoryRes{Id: id}, nil
}

func (s *Service) UpdateRepository(ctx context.Context, req *api.Repository) (*api.Empty, error) {
	err := s.storage.UpdateRepository(ctx, &storage.Repository{
		ID:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Endpoint: req.Endpoint,
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &api.Empty{}, nil
}
