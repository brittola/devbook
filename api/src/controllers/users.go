package controllers

import "net/http"

// CreateUser adiciona um usuário ao banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

// GetUsers busca os usuários no banco de dados
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando usuários"))
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
