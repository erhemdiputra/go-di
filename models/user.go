package models

type UserCookie struct {
	Name string `json:"name"`
}

type User struct {
	ID       int64  `db:"id" json:"id"`
	FullName string `db:"fullname" json:"fullname"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}
