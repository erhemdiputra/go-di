package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/erhemdiputra/go-di/models"
)

type IUserRepo interface {
	IsValidUser(ctx context.Context, name string, password string) (*models.User, error)
}

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) IUserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) IsValidUser(ctx context.Context, name string, password string) (*models.User, error) {
	query := `SELECT id, fullname, username, password FROM users WHERE username = ? AND password = ?`

	var user models.User
	err := repo.DB.QueryRowContext(ctx, query, name, password).Scan(&user.ID, &user.FullName,
		&user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("database query user error : {%+v}", err)
	}

	return &user, nil
}
