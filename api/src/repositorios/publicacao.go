package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Publicacao struct {
	db *sql.DB
}

func NovoRepositorioPublicacao(db *sql.DB) *Publicacao {
	return &Publicacao{db}
}

func (repositorio Publicacao) BuscarPublicacao(publicacaoId uint64) (modelos.Publicacao, error) {

	fmt.Println(publicacaoId)
	linhas, erro := repositorio.db.Query(`select p.id, p.titulo, p.conteudo, p.autor_id, u.nick, p.curtidas, p.criadoem from publicacoes p 
		inner join usuarios u on u.id = p.autor_id 
		where p.id = ?`, publicacaoId)

	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linhas.Close()

	var publicacao modelos.Publicacao
	if linhas.Next() {
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.AutorNick,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}
	fmt.Println(publicacao)

	return publicacao, nil

}
