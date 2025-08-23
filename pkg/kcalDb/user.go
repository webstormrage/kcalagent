package kcaldb

import (
	"ai-kcal-agent/pkg/appContext"
	"database/sql"
)

type User struct {
	ID           int64
	Login        string
	PasswordHash string
}

func GetUserByLogin(login string) (*User, error) {
	ctx := appContext.Get()
	db, err := sql.Open("postgres", ctx.DataSourceName)
	defer db.Close()
	if err != nil {
		ctx.Logger.Println("[DB]: connection error")
		return nil, err
	}
	ctx.Logger.Printf("[DB]: get user by login query <%s>\n", login)
	var u User
	err = db.QueryRow(
		"SELECT id, login, password_hash FROM users WHERE login = $1 LIMIT 1",
		login,
	).Scan(&u.ID, &u.Login, &u.PasswordHash)
	if err != nil {
		ctx.Logger.Printf("[DB]: get user by login query failed <%s>\n", login)
		return nil, err
	}
	ctx.Logger.Printf("[DB]: get user by login query success<%s>\n", login)
	return &u, nil
}
