package mysql

import (
	"database/sql"
	"time"

	"merlin.com/box/pkg/models"
)

 
func (m *TravelModel) InsertComent(IdComentarios int, Texto string, DataComent time.Timer, Upvote int, Downvote int, U_IdUsuario int, H_IdHotel int  )  (int, error) {
	stmt := `INSERT INTO comentarios ( idcomentario, Texto, DataComentario, UpVotes, DownVotes, QuantVotes, Usuarios_idusuario, Hoteis_idhoteis) 
            VALUES(?,?,?,?,?,?,?,?)`

	result, err := m.DB.Exec(stmt, IdComentarios, Texto, DataComent, Upvote, Downvote,U_IdUsuario,H_IdHotel)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *TravelModel) GetComent(id int) (*models.Comentarios, error) { //te, algum erro
	stmt := `SELECT idcomentario, Texto, DataComentario, UpVotes, DownVotes, Usuarios_idusuario, Hoteis_idhoteis FROM comentarios
           WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	s := &models.Comentarios{}

	err := row.Scan(&s.IdComentarios, &s.Texto, &s.DataComent, &s.Upvote, &s.Downvote, &s.U_IdUsuario,&s.H_IdHotel)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}
//volta depois

func (m *TravelModel) LatestComent() ([]*models.Comentarios, error) {
	stmt := `SELECT * FROM comentarios INNER JOIN Hoteis
           on comentarios.Hoteis_IdHoteis = Hoteis.IdHoteis ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	ComentariosS := []*models.Comentarios{}
	for rows.Next() {
    s:= &models.Comentarios{}
    err = rows.Scan(&s.IdComentarios, &s.Texto, &s.DataComent, &s.Upvote, &s.Downvote, &s.U_IdUsuario,&s.H_IdHotel)
    if err != nil{
      return nil, err
    }
    ComentariosS = append(ComentariosS, s)
	
  }
  err = rows.Err()
  if err != nil{ //executa funcao, coloca na variavel 
    return nil, err //executa
  }
	return nil, nil
}