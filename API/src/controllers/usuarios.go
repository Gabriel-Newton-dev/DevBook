package controllers

import "net/http"

//CriarUsuario irá dar um post - criar usuário.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!"))
}

// BuscarUsuarios irá buscar todos os usuários salvos no database.
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuário!"))
}

//Buscarusuario irá buscar um usuário pelo ID.
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca usuário pelo Id!"))
}

// Atualizarusuario irá atualizar as informações de usuárioID no database.
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário!"))
}

// DeletarUsuario irá deletar um usuárioID no banco.
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário!"))
}
