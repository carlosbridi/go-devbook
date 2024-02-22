package modelos

import (
	"api/src/modelos"
	"errors"
	"strings"
	"time"
)

// Modelo de publicação feita por um usuário
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorId   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"` // publicacao sem curtida não iriam ter no json de saída, e como o valor padrão é zero, pode ficar sem
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

func (publicacao *modelos.Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

func (publicacao *modelos.Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("Título é obrigatório")
	}
	if publicacao.Conteudo == "" {
		return errors.New("Conteúdo é obrigatório")
	}
	return nil
}

func (publicacao *modelos.Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
