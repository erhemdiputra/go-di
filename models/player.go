package models

type PlayerForm struct {
	FullName string `json:"full_name,omitempty"`
	Club     string `json:"club,omitempty"`
}

type PlayerResponse struct {
	ID       int    `db:"id" json:"id"`
	FullName string `db:"full_name" json:"full_name"`
	Club     string `db:"club" json:"club"`
}
