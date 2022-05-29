package mysql

import (
	"database/sql"
	"merlin.com/box/pkg/models"
)

 
func (m *TravelModel) InsertHoteis(IdHoteis int, Nome string, Cidade string, Pais string,Cnpj string)  (int, error) {
	stmt := `INSERT INTO hoteis (idHoteis, Nome, Email, Cpf, Telefone) 
            VALUES(?,?,?,?,?,?)`

	result, err := m.DB.Exec(stmt, IdHoteis, Nome, Cidade, Pais, Cnpj)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *TravelModel) GetHoteis(id int) (*models.Hoteis, error) { //te, algum erro
	stmt := `SELECT idHoteis, Nome, Email, Cpf, Telefone FROM hoteis
           WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	s := &models.Hoteis{}

	err := row.Scan(&s.IdHoteis, &s.Nome, &s.Cidade, &s.Pais, &s.Cnpj)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}
//volta depois

func (m *TravelModel) MostPopularHoteis() ([]*models.Hoteis, error) {
	stmt := `SELECT idHoteis, Nome, UpVotes, DownVotes, FROM hoteis
          ORDER BY created DESC LIMIT 10 `

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	HoteisS := []*models.Hoteis{}
	for rows.Next() {
    s:= &models.Hoteis{}
    err = rows.Scan(&s.IdHoteis, &s.Nome, &s.Cidade, &s.Pais, &s.Cnpj)
    if err != nil{
      return nil, err
    }
    HoteisS = append(HoteisS, s)
	
  }
  err = rows.Err()
  if err != nil{ //executa funcao, coloca na variavel 
    return nil, err //executa
  }
	return nil, nil
}

func (m *TravelModel) MostViewsHoteis() ([]*models.Hoteis, error) {
	stmt := `SELECT idHoteis, Nome, UpVotes, DownVotes, FROM hoteis
          WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10 `

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	HoteisS := []*models.Hoteis{}
	for rows.Next() {
    s:= &models.Hoteis{}
    err = rows.Scan(&s.IdHoteis, &s.Nome, &s.UpVotes, &s.DownVotes,) 
     
    if err != nil{
      return nil, err
    }
    HoteisS = append(HoteisS, s)
	
  }
  err = rows.Err()
  if err != nil{ //executa funcao, coloca na variavel 
    return nil, err //executa
  }
	return nil, nil
}
func (m *TravelModel) MostPopularHoteis() ([]*models.Hoteis, error) {
	stmt := `SELECT idHoteis, Nome, UpVotes, DownVotes, FROM hoteis
          ORDER BY created DESC LIMIT 10 `

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	HoteisS := []*models.Hoteis{}
	for rows.Next() {
    s:= &models.Hoteis{}
    err = rows.Scan(&s.IdHoteis, &s.Nome, &s.Cidade, &s.Pais, &s.Cnpj)
    if err != nil{
      return nil, err
    }
    HoteisS = append(HoteisS, s)
	
  }
  err = rows.Err()
  if err != nil{ //executa funcao, coloca na variavel 
    return nil, err //executa
  }
	return nil, nil
}

func (m *TravelModel) TopViewsHoteis() ([]*models.Hoteis, error) {
	stmt := `SELECT idHoteis, Nome, UpVotes, DownVotes, Views FROM hoteis ORDER BY Views DESC`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	HoteisS := []*models.Hoteis{}
	for rows.Next() {
    s:= &models.Hoteis{}
    err = rows.Scan(&s.IdHoteis, &s.Nome, &s.UpVotes, &s.DownVotes,) 
     
    if err != nil{
      return nil, err
    }
    HoteisS = append(HoteisS, s)
	
  }
  err = rows.Err()
  if err != nil{ //executa funcao, coloca na variavel 
    return nil, err //executa
  }
	return nil, nil
}