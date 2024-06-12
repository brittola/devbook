package controllers

import (
	"devbook-api/src/database"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// CreateUser adiciona um usuário ao banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	usersRepository := repositories.NewUsersRepository(db)
	user.ID, err = usersRepository.Create(user)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// GetUsers busca os usuários no banco de dados
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrUsername := strings.ToLower(r.URL.Query().Get("user"))

	if nameOrUsername == "" {
		responses.Error(w, http.StatusBadRequest, errors.New("a query 'user' não pode ser vazia"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	usersRepository := repositories.NewUsersRepository(db)
	users, err := usersRepository.Search(nameOrUsername)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)

}

// GetUser busca um usuário no banco de dados por ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando usuário"))
}

// UpdateUser busca um usuário no banco de dados por ID e atualiza informações
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

// DeleteUser exclui um usuário do banco de dados por ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
