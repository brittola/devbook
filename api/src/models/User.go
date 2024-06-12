package models

import (
	"errors"
	"strings"
	"time"
)

// User representa um usuário do sistema Devbook
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare valida e formata os campos da struct User
func (u *User) Prepare() error {
	if err := u.validate(); err != nil {
		return err
	}
	u.format()

	return nil
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}
	if u.Username == "" {
		return errors.New("o username é obrigatório e não pode estar em branco")
	}
	if u.Email == "" {
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}
	if u.Password == "" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)
}
