package models

import (
	"database/sql"
)

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type userModel struct {
	db *sql.DB
}

func (m *userModel) GetUserByID(id int) (*User, error) {
	row := m.db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err // Handle other database errors
	}
	return user, nil
}

func (m *userModel) GetAllUsers() ([]User, error) {
	rows, err := m.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (m *userModel) CreateUser(user *User) error {
	stmt, err := m.db.Prepare("INSERT INTO users (id, username, email, password_hash) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Username, user.Email, user.PasswordHash)
	return err
}

func (m *userModel) UpdateUser(user *User) error {
	stmt, err := m.db.Prepare("UPDATE users SET username = ?, email = ?, password_hash = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.PasswordHash, user.ID)
	return err
}

func (m *userModel) DeleteUser(id int) error {
	stmt, err := m.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
