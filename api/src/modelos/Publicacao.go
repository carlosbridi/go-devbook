package modelos

import "time"

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
