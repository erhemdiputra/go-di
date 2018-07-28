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

func (c *PlayerController) GetList(ctx context.Context, form models.PlayerForm) ([]models.PlayerResponse, error) {
	return c.PlayerService.GetList(ctx, form)
}
