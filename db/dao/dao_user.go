package dao

import (
	"crypto/sha256"
	"fiap/ancora/db"
	"fiap/ancora/model"
	"fmt"
)

/* GetAllUsers returns all users */
func GetAllUsers() []model.User {
	sql := "SELECT id, username, password FROM users"
	rows, err := db.GetDB().Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}
	return users
}

func GetUserByUsername(username string) model.User {
	sql := "SELECT id, username, password FROM users WHERE username = $1"
	row := db.GetDB().QueryRow(sql, username)
	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		fmt.Println(err)
	}
	return user
}

func InsertUser(user model.User) {
	passwordHash := sha256.New().Sum([]byte(user.Password))
	sql := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := db.GetDB().Exec(sql, user.Username, passwordHash)
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateUser(user model.User) {
	passwordHash := sha256.New().Sum([]byte(user.Password))
	sql := "UPDATE users SET username = $1, password = $2 WHERE id = $3"
	_, err := db.GetDB().Exec(sql, user.Username, passwordHash, user.ID)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteUser(id int) {
	sql := "DELETE FROM users WHERE id = $1"
	_, err := db.GetDB().Exec(sql, id)
	if err != nil {
		fmt.Println(err)
	}
}
