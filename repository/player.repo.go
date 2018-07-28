package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/erhemdiputra/go-di/models"
)

type IPlayerRepo interface {
	GetList(ctx context.Context) ([]models.Player, error)
}

type PlayerRepo struct {
	DB *sql.DB
}

func NewPlayerRepo(db *sql.DB) IPlayerRepo {
	return &PlayerRepo{
		DB: db,
	}
}

func (repo *PlayerRepo) GetList(ctx context.Context) ([]models.Player, error) {
	query := `SELECT * FROM players`

	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("database query player list error : {%+v}", err)
	}
	defer rows.Close()

	playerList := []models.Player{}

	for rows.Next() {
		var player models.Player

		err = rows.Scan(&player.ID, &player.FullName, &player.Club)
		if err != nil {
			return nil, fmt.Errorf("database scan rows error : {%+v}", err)
		}

		playerList = append(playerList, player)
	}

	return playerList, nil
}
