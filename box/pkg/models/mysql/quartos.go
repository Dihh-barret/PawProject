package mysql

import (
	"database/sql"
	"merlin.com/box/pkg/models"
)
 
func (m *TravelModel) InsertQuartos( IdQuartos int, Numero int , Upvote int, Downvote int, Preço float32, Status int, H_IdHotel int )  (int, error) {
	stmt := `INSERT INTO quartos (idquartos, NumeroQuarto, UpVotes, DownVotes, QuantVotes, Preco, Hoteis_idhoteis) 
            VALUES(?,?,?,?,?,?)`

	result, err := m.DB.Exec(stmt, IdQuartos, Numero, Upvote, Downvote, Preço,Status,H_IdHotel)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *TravelModel) GetQuatos(id int) (*models.Quartos, error) { //te, algum erro
	stmt := `SELECT idquartos, NumeroQuarto, UpVotes, DownVotes, QuantVotes, Preco, Hoteis_idhoteis FROM quartos
           WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	s := &models.Quartos{}

	err := row.Scan(&s.IdQuartos, &s.Numero, &s.Upvote, &s.Downvote, &s.Preço, &s.Status, &s.H_IdHotel)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}

func (m *TravelModel) LatestQuartos() ([]*models.Quartos, error) {
	stmt := `SELECT idquartos, NumeroQuarto, UpVotes, DownVotes, QuantVotes, Preco, Hoteis_idhoteis FROM quartos
           WHERE id = ? `

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	QuartosS := []*models.Quartos{}
	for rows.Next() {
    s:= &models.Quartos{}
    err = rows.Scan(&s.IdQuartos, &s.Numero, &s.Upvote, &s.Downvote, &s.Preço, &s.Status, &s.H_IdHotel)
    if err != nil{
      return nil, err
    }
    QuartosS = append(QuartosS, s)
	
  }
  err = rows.Err()
  if err != nil{ //executa funcao, coloca na variavel 
    return nil, err //executa
  }
	return nil, nil
}

