package mysql

import (
	"database/sql"
	"merlin.com/box/pkg/models"
)
 
func (m *TravelModel) InsertQuartos( IdQuartos int, Numero int , Preço float32,H_IdHotel int,Descricao string, Wifi bool, Ar_Cond bool)  (int, error) {
	stmt := `INSERT INTO quartos (idquartos, NumeroQuarto, Preco, Hoteis_idhoteis, Descricao, Wi fi, Ar Condicionado) 
            VALUES(?,?,?,?,?,?)`

	result, err := m.DB.Exec(stmt, IdQuartos, Numero, Preço,H_IdHotel,Descricao, Wifi, Ar_Cond)
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
	stmt := `SELECT idquartos, NumeroQuarto, Preco, Hoteis_idhoteis, Descricao, Wi fi, Ar Condicionado WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	s := &models.Quartos{}

	err := row.Scan(&s.IdQuartos, &s.Numero, &s.Preço, &s.H_IdHotel, &s.Descricao, &s.Wifi, &s.ArCond)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}

func (m *TravelModel) BQuartos() ([]*models.Quartos, error) {
	stmt := `SELECT idquartos, NumeroQuarto, Preco, Hoteis_idhoteis, Descricao, Wi fi, Ar Condicionado WHERE id = ? `

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	QuartosS := []*models.Quartos{}
	for rows.Next() {
    s:= &models.Quartos{}
    err = rows.Scan(&s.IdQuartos, &s.Numero, &s.Preço, &s.H_IdHotel, &s.Descricao, &s.Wifi, &s.ArCond)
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

