package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/erhemdiputra/go-di/models"
)

type IPlayerRepo interface {
	GetList(ctx context.Context, form models.PlayerForm) ([]models.PlayerResponse, error)
}

type PlayerRepo struct {
	DB *sql.DB
}

func NewPlayerRepo(db *sql.DB) IPlayerRepo {
	return &PlayerRepo{
		DB: db,
	}
}

func (repo *PlayerRepo) GetList(ctx context.Context, form models.PlayerForm) ([]models.PlayerResponse, error) {
	query := repo.BuildQueryPlayerList(form)
	log.Printf("[PlayerRepo] -> Get List Query : %s\n", query)

	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("database query player list error : {%+v}", err)
	}
	defer rows.Close()

	playerList := []models.PlayerResponse{}

	for rows.Next() {
		var player models.PlayerResponse

		err = rows.Scan(&player.ID, &player.FullName, &player.Club)
		if err != nil {
			return nil, fmt.Errorf("database scan rows error : {%+v}", err)
		}

		playerList = append(playerList, player)
	}

	return playerList, nil
}
