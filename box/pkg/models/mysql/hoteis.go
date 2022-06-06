package mysql

import (
	"database/sql"
	"merlin.com/box/pkg/models"
)

 
func (m *TravelModel) InsertHoteis(IdHoteis int, Nome string, Cidade string, Pais string,Cnpj string,S_idSenhas int, UpVotes int,DownVotes int,HospNum int )  (int, error) {
	stmt := `INSERT INTO hoteis (idhoteis,Nome, CNPJ, Cidade, Pais, Senhas_idSenhas, UpVotes, DownVotes, HospNum) 
            VALUES(?,?,?,?,?,?,?,?,?)`

	result, err := m.DB.Exec(stmt, IdHoteis, Nome,Cnpj, Cidade, Pais, S_idSenhas,UpVotes,DownVotes,HospNum)
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
	stmt := `SELECT idhoteis, Nome,CNPJ, Cidade,Pais, Pais, Senhas_idSenhas, UpVotes, DownVotes, HospNum  FROM hoteis
           WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	s := &models.Hoteis{}

	err := row.Scan(&s.IdHoteis, &s.Nome, &s.Cidade, &s.Pais, &s.Cnpj, &s.Cidade, &s.Pais, &s.Cnpj)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}
//volta depois

func (m *TravelModel) MostPopularHoteis() ([]*models.Hoteis, error) {
	stmt := `SELECT idhoteis, Nome,Pais, UpVotes, DownVotes, HospNum FROM hoteis`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	HoteisS := []*models.Hoteis{}
	for rows.Next() {
    s:= &models.Hoteis{}
    err = rows.Scan(&s.IdHoteis, &s.Nome, &s.Pais, &s.UpVotes, &s.DownVotes, &s.HospNum)
    if err != nil{
      return nil, err
    }
    HoteisS = append(HoteisS, s)
	
  }
  err = rows.Err()
  if err != nil{ //executa funcao, coloca na variavel 
    return nil, err //executa
  }
	return HoteisS, nil
}
/*func (m *TravelModel) MostPopularHoteis() ([]*models.Hoteis,[]*models.Quartos, error,error) {
	stmt := `SELECT hoteis.idhoteis, hoteis.Nome,hoteis.Pais, hoteis.UpVotes, hoteis.DownVotes, hoteis.HospNum FROM hoteis INNER JOIN quartos ON hoteis.idhoteis= quartos.Hoteis_idhoteis
          ORDER BY UpVotes DESC LIMIT 10 `

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil,nil, err
	}
	defer rows.Close()

	//crando slice
	HoteisS ,QuartosS := []*models.Hoteis{},[]*models.Quartos{}
	for rows.Next() {
    s:= &models.Hoteis{}
    err = rows.Scan(&s.IdHoteis, &s.Nome, &s.Pais, &s.UpVotes, &s.DownVotes)
    if err != nil{
      return nil,nil, err,err2
    }
    HoteisS = append(HoteisS, s)
    QuartosS = append(QuartosS, s)
	
  }
  err = rows.Err()
  if err != nil{ //executa funcao, coloca na variavel 
    return nil,nil, err //executa
  }
	return nil,nil, nil
}*/
func (m *TravelModel) TestHoteis() ([]*models.Hoteis, error) {
	stmt := `SELECT idhoteis,Nome, CNPJ, Cidade, Pais, Senhas_idSenhas, UpVotes, DownVotes, HospNum FROM hoteis `

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	HoteisS := []*models.Hoteis{}
	for rows.Next() {
    s:= &models.Hoteis{}
    err = rows.Scan(&s.IdHoteis, &s.Nome, &s.Cnpj, &s.Pais, &s.Cidade, &s.S_IdSenhas, &s.UpVotes, &s.DownVotes, &s.HospNum)
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

// func (m *TravelModel) MostViewsHoteis() ([]*models.Hoteis, error) {
// 	stmt := `SELECT idHoteis, Nome, UpVotes, DownVotes, FROM hoteis
//           WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10 `

// 	rows, err := m.DB.Query(stmt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	//crando slice
// 	HoteisS := []*models.Hoteis{}
// 	for rows.Next() {
//     s:= &models.Hoteis{}
//     err = rows.Scan(&s.IdHoteis, &s.Nome, &s.UpVotes, &s.DownVotes) 
     
//     if err != nil{
//       return nil, err
//     }
//     HoteisS = append(HoteisS, s)
	
//   }
//   err = rows.Err()
//   if err != nil{ //executa funcao, coloca na variavel 
//     return nil, err //executa
//   }
// 	return nil, nil
// }
// func (m *TravelModel) MostPopularHoteis() ([]*models.Hoteis, error) {
// 	stmt := `SELECT idHoteis, Nome, UpVotes, DownVotes, FROM hoteis
//           ORDER BY created DESC LIMIT 10 `

// 	rows, err := m.DB.Query(stmt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	//crando slice
// 	HoteisS := []*models.Hoteis{}
// 	for rows.Next() {
//     s:= &models.Hoteis{}
//     err = rows.Scan(&s.IdHoteis, &s.Nome, &s.Cidade, &s.Pais, &s.Cnpj)
//     if err != nil{
//       return nil, err
//     }
//     HoteisS = append(HoteisS, s)
	
//   }
//   err = rows.Err()
//   if err != nil{ //executa funcao, coloca na variavel 
//     return nil, err //executa
//   }
// 	return nil, nil
// }

func (m *TravelModel) TopViewsHoteis() ([]*models.Hoteis, error) {
	stmt := `SELECT * FROM hoteis ORDER BY Views DESC`

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