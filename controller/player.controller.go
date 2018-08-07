package controller

import (
	"context"

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

func (c *PlayerController) GetList(ctx context.Context, form models.PlayerForm) ([]models.Player, error) {
	return c.PlayerService.GetList(ctx, form)
}

func (c *PlayerController) Add(ctx context.Context, form models.PlayerForm) (int64, error) {
	return c.PlayerService.Add(ctx, form)
}

func (c *PlayerController) GetByID(ctx context.Context, id int64) (*models.Player, error) {
	return c.PlayerService.GetByID(ctx, id)
}

func (c *PlayerController) Update(ctx context.Context, id int64, form models.PlayerForm) (int64, error) {
	return c.PlayerService.Update(ctx, id, form)
}
