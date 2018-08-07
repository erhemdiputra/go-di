package models

import "html"

type PlayerForm struct {
	FullName string `json:"full_name,omitempty"`
	Club     string `json:"club,omitempty"`
}

type PlayerResponse struct {
	ID       int64  `db:"id" json:"id"`
	FullName string `db:"full_name" json:"full_name"`
	Club     string `db:"club" json:"club"`
}

func (form *PlayerForm) Sanitize() {
	form.FullName = html.EscapeString(form.FullName)
	form.Club = html.EscapeString(form.Club)
}

func (form *PlayerForm) IsEmpty() bool {
	return *form == (PlayerForm{}) || form.FullName == ""
}
