package service

import (
	"context"
	"fmt"

	infraMemCache "github.com/erhemdiputra/go-di/infrastructure_services/memcache"
	"github.com/erhemdiputra/go-di/models"
	"github.com/erhemdiputra/go-di/repository"
)

type IPlayerService interface {
	GetList(ctx context.Context, form models.PlayerForm) ([]models.Player, error)
	Add(ctx context.Context, form models.PlayerForm) (int64, error)
	GetByID(ctx context.Context, id int64) (models.Player, error)
	Update(ctx context.Context, id int64, form models.PlayerForm) (int64, error)
}

type PlayerService struct {
	PlayerRepo repository.IPlayerRepo
	MemCache   infraMemCache.IMemCache
}

func NewPlayerService(playerRepo repository.IPlayerRepo, memCache infraMemCache.IMemCache) IPlayerService {
	return &PlayerService{
		PlayerRepo: playerRepo,
		MemCache:   memCache,
	}
}

func (s *PlayerService) GetList(ctx context.Context, form models.PlayerForm) ([]models.Player, error) {
	var playerList []models.Player

	_ = s.MemCache.GetCacheTTLJSON(infraMemCache.Key15Min, models.KeyCachePlayerList, &playerList)

	if len(playerList) > 0 {
		return playerList, nil
	}

	playerList, err := s.PlayerRepo.GetList(ctx, form)
	if err != nil {
		return nil, err
	}

	_ = s.MemCache.SetCacheTTLJSON(infraMemCache.Key15Min, models.KeyCachePlayerList, playerList)

	return playerList, nil
}

func (s *PlayerService) Add(ctx context.Context, form models.PlayerForm) (int64, error) {
	id, err := s.PlayerRepo.Add(ctx, form)
	if err != nil {
		return 0, err
	}

	if err := s.MemCache.DeleteCache(infraMemCache.Key15Min, models.KeyCachePlayerList); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *PlayerService) GetByID(ctx context.Context, id int64) (models.Player, error) {
	var player models.Player

	key := fmt.Sprintf(models.KeyCachePlayerID, id)

	_ = s.MemCache.GetCacheTTLJSON(infraMemCache.Key15Min, key, &player)

	if player.ID != 0 {
		return player, nil
	}

	player, err := s.PlayerRepo.GetByID(ctx, id)
	if err != nil {
		return models.Player{}, err
	}

	_ = s.MemCache.SetCacheTTLJSON(infraMemCache.Key15Min, key, player)

	return player, nil
}

func (s *PlayerService) Update(ctx context.Context, id int64, form models.PlayerForm) (int64, error) {
	id, err := s.PlayerRepo.Update(ctx, id, form)
	if err != nil {
		return 0, err
	}

	if err := s.MemCache.DeleteCache(infraMemCache.Key15Min, models.KeyCachePlayerList); err != nil {
		return 0, err
	}

	return id, nil
}
