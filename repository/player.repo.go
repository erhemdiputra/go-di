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
	Add(ctx context.Context, form models.PlayerForm) (int64, error)
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
	query := repo.BuildQueryGetList(form)
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

func (repo *PlayerRepo) Add(ctx context.Context, form models.PlayerForm) (int64, error) {
	query := `INSERT INTO players(full_name, club) VALUES (?, ?)`
	log.Printf("[PlayerRepo] -> Insert : %s\n", query)

	res, err := repo.DB.ExecContext(ctx, query, form.FullName, form.Club)
	if err != nil {
		return 0, fmt.Errorf("database insert player error : {%+v}", err)
	}

	return res.LastInsertId()
}
