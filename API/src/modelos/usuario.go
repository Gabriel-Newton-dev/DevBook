package modelos

import "time"

// Usuario representa um usuádio utilizando a rede social.
type Usuario struct {
	ID        uint64    `json:"ïd,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Senha     string    `json:"senha,omitempty"`
	CriandoEm time.Time `json:"CriandoEm,omitempty"`
}
