package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	// Iremos criar o CRUD usando o Struct que contruímos na rotas.go(Create, Read, Update e Delete)

	{URI: "/usuarios",
		Metodo:             http.MethodPost, // cadastra um usuário
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{URI: "/usuarios",
		Metodo:             http.MethodGet, // busca todos os usuários.
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{URI: "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet, // busca os usuários pelo ID
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{URI: "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut, // put - atualizar usuário
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{URI: "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete, // Delete o usuário pelo ID
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
