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

func (c *PlayerController) GetList(ctx context.Context) ([]models.Player, error) {
	return c.PlayerService.GetList(ctx)
}
