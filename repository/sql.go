package repository

import (
	"fmt"

	"github.com/erhemdiputra/go-di/models"
)

func (repo *PlayerRepo) BuildQueryGetList(form models.PlayerForm) string {
	query := `SELECT id, full_name, club FROM players`

	if form == (models.PlayerForm{}) {
		return query
	}

	query += ` WHERE`
	var needAnd bool

	if form.FullName != "" {
		query += fmt.Sprintf(` full_name LIKE '%%%s%%'`, form.FullName)
		needAnd = true
	}

	if form.Club != "" {
		if needAnd {
			query += fmt.Sprintf(` AND`)
		}

		query += fmt.Sprintf(` club LIKE '%%%s%%'`, form.Club)
		needAnd = true
	}

	return query
}
