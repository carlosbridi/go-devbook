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

	statement, erro := repositorio.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values (?,?,?)")

	if erro != nil {
		return 0, erro
	}

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorId)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Publicacao) BuscarPublicacoes(autorId uint64) ([]modelos.Publicacao, error) {

	linhas, erro := repositorio.db.Query(
		`select distinct p.*, u.nick 
			from publicacoes p 
		  inner join usuarios u on p.autor_id = u.id
		  inner join seguidores s on s.usuario_id = p.id
	   	  where u.id = ? or s.seguidor_id = ?`, autorId, autorId)

	if erro != nil {
		return nil, erro
	}

	var publicacaoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.Curtidas,
			&publicacao.AutorId,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacaoes = append(publicacaoes, publicacao)
	}
	return publicacaoes, nil

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

func (repositorio Publicacao) DeletarPublicacao(autorId, publicacaoId uint64) error {

	statement, erro := repositorio.db.Prepare("DELETE FROM publicacaoes where id = ? where autor_id")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId, autorId); erro != nil {
		return erro
	}

	return nil

}
