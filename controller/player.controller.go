package controller

import (
	"context"
	"errors"

	"github.com/erhemdiputra/go-di/models"
	"github.com/erhemdiputra/go-di/service"
)

type PlayerController struct {
	PlayerService service.IPlayerService
}

func NewPlayerController(PlayerService service.IPlayerService) *PlayerController {
	return &PlayerController{
		PlayerService: PlayerService,
	}
}

func (c *PlayerController) GetList(ctx context.Context, form models.PlayerForm) ([]models.PlayerResponse, error) {
	return c.PlayerService.GetList(ctx, form)
}

func (c *PlayerController) Add(ctx context.Context, form models.PlayerForm) (int64, error) {
	form.Sanitize()
	return c.PlayerService.Add(ctx, form)
}

func (c *PlayerController) GetByID(ctx context.Context, id int64) (*models.PlayerResponse, error) {
	if id <= 0 {
		return nil, errors.New("Invalid Player ID")
	}
	return c.PlayerService.GetByID(ctx, id)
}
