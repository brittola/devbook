package repositories

import (
	"database/sql"
	"devbook-api/src/models"
	"fmt"
)

type users struct {
	db *sql.DB
}

// NewUsersRepository retorna um repositório de usuários, para manipular a tabela de usuários no banco de dados
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// Create insere um usuário no banco de dados
func (repo users) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare(
		"INSERT INTO users (name, username, email, password) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}

// Search retorna um slice de usuários que atendem a um filtro de nome ou username
func (repo users) Search(nameOrUsername string) ([]models.User, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername)

	rows, err := repo.db.Query(
		"SELECT id, name, username, email, created_at FROM users WHERE name LIKE ? OR username LIKE ?",
		nameOrUsername, nameOrUsername,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
