package service

import (
	"context"

	"github.com/erhemdiputra/go-di/models"
	"github.com/erhemdiputra/go-di/repository"
)

type IPlayerService interface {
	GetList(ctx context.Context, form models.PlayerForm) ([]models.PlayerResponse, error)
	Add(ctx context.Context, form models.PlayerForm) (int64, error)
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
