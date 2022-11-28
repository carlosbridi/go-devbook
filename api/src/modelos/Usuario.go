package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuário
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("Nome é obrigatório")
	}

	if usuario.Nick == "" {
		return errors.New("Nick é obrigatório")
	}

	if usuario.Email == "" {
		return errors.New("Email é obrigatório")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("Senha é obrigatório")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Senha = strings.TrimSpace(usuario.Senha)
	usuario.Email = strings.TrimSpace(usuario.Email)

}
