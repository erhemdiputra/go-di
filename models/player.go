package models

type Player struct {
	ID       int    `db:"id" json:"id"`
	FullName string `db:"full_name" json:"full_name"`
	Club     string `db:"club" json:"club"`
}
