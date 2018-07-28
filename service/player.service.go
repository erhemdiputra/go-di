package service

import (
	"context"

	"github.com/erhemdiputra/go-di/models"
	"github.com/erhemdiputra/go-di/repository"
)

type IPlayerService interface {
	GetList(ctx context.Context) ([]models.Player, error)
}

type PlayerService struct {
	PlayerRepo repository.IPlayerRepo
}

func NewPlayerService(playerRepo repository.IPlayerRepo) IPlayerService {
	return &PlayerService{
		PlayerRepo: playerRepo,
	}
}

func (s *PlayerService) GetList(ctx context.Context) ([]models.Player, error) {
	return s.PlayerRepo.GetList(ctx)
}
