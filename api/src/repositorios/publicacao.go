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

func (repositorio Publicacao) CriarPublicacao(publicacao modelos.Publicacao) (uint64, error) {

	statement, erro := repositorio.db.Prepare("insert into publicacao (titluo, conteudo, autor_id, autor_nick) values (?,?,?,?,?)")

	if erro != nil {
		return 0, erro
	}

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorId, publicacao.AutorNick)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Publicacao) BuscarPublicacao(publicacaoId uint64) (modelos.Publicacao, error) {

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
