package controllers

import "net/http"

//CriarUsuario irá dar um post - criar usuário.
func CriarUsuario(W http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!"))
}

// BuscarUsuarios irá buscar todos os usuários salvos no database.
func BuscarUsuarios(W http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuário!"))
}

//Buscarusuario irá buscar um usuário pelo ID.
func BuscarUsuario(W http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca usuário pelo Id!"))
}

// Atualizarusuario irá atualizar as informações de usuárioID no database.
func AtualizarUsuario(W http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário!"))
}

// DeletarUsuario irá deletar um usuárioID no banco.
func DeletarUsuario(W http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário!"))
}
