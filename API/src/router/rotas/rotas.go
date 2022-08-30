package rotas

import "net/http"

// Rota representa todas as rotas da API, serão contruídas seguindo como base essa Struct.
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}
