package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/erhemdiputra/go-di/models"
)

type IPlayerRepo interface {
	GetList(ctx context.Context, form models.PlayerForm) ([]models.Player, error)
	Add(ctx context.Context, form models.PlayerForm) (int64, error)
	GetByID(ctx context.Context, id int64) (models.Player, error)
	Update(ctx context.Context, id int64, form models.PlayerForm) (int64, error)
}

type PlayerRepo struct {
	DB *sql.DB
}

func NewPlayerRepo(db *sql.DB) IPlayerRepo {
	return &PlayerRepo{
		DB: db,
	}
}

func (repo *PlayerRepo) GetList(ctx context.Context, form models.PlayerForm) ([]models.Player, error) {
	query := repo.BuildQueryGetList(form)

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

func (repo *PlayerRepo) Add(ctx context.Context, form models.PlayerForm) (int64, error) {
	query := `INSERT INTO players(full_name, club) VALUES (?, ?)`

	res, err := repo.DB.ExecContext(ctx, query, form.FullName, form.Club)
	if err != nil {
		return 0, fmt.Errorf("database insert player error : {%+v}", err)
	}

	return res.LastInsertId()
}

func (repo *PlayerRepo) GetByID(ctx context.Context, id int64) (models.Player, error) {
	query := `SELECT id, full_name, club FROM players WHERE id = ?`

	var player models.Player

	err := repo.DB.QueryRowContext(ctx, query, id).Scan(&player.ID, &player.FullName, &player.Club)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Player{}, err
		}
		return models.Player{}, fmt.Errorf("database query player by id error : {%+v}", err)
	}

	return player, nil
}

func (repo *PlayerRepo) Update(ctx context.Context, id int64, form models.PlayerForm) (int64, error) {
	query := `UPDATE players SET full_name = ?, club = ? WHERE id = ?`

	res, err := repo.DB.ExecContext(ctx, query, form.FullName, form.Club, id)
	if err != nil {
		return 0, fmt.Errorf("database update player error : {%+v}", err)
	}

	return res.RowsAffected()
}
