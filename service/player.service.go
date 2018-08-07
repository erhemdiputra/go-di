package service

import (
	"context"

	"github.com/erhemdiputra/go-di/models"
	"github.com/erhemdiputra/go-di/repository"
)

type IPlayerService interface {
	GetList(ctx context.Context, form models.PlayerForm) ([]models.PlayerResponse, error)
	Add(ctx context.Context, form models.PlayerForm) (int64, error)
	GetByID(ctx context.Context, id int64) (*models.PlayerResponse, error)
}

type PlayerService struct {
	PlayerRepo repository.IPlayerRepo
}

func NewPlayerService(playerRepo repository.IPlayerRepo) IPlayerService {
	return &PlayerService{
		PlayerRepo: playerRepo,
	}
}

func (s *PlayerService) GetList(ctx context.Context, form models.PlayerForm) ([]models.PlayerResponse, error) {
	return s.PlayerRepo.GetList(ctx, form)
}

func (s *PlayerService) Add(ctx context.Context, form models.PlayerForm) (int64, error) {
	return s.PlayerRepo.Add(ctx, form)
}

func (s *PlayerService) GetByID(ctx context.Context, id int64) (*models.PlayerResponse, error) {
	return s.PlayerRepo.GetByID(ctx, id)
}
