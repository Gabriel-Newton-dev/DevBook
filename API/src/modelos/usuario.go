package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuádio utilizando a rede social.
type Usuario struct {
	ID        uint64    `json:"ïd,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Senha     string    `json:"senha,omitempty"`
	CriandoEm time.Time `json:"CriandoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido.
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O Nick é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O E-mail é obrigatório e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome) // TrimSpaca é para tirar os espaços nas extremidades.
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
