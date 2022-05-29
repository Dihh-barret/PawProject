package mysql

import (
	"database/sql"
	"merlin.com/box/pkg/models"
  "time"
)


func (m *TravelModel) InsertCartoes(IdCartoes int, Numero string,Cvv int, DataEmi time.Timer,DataVal time.Timer, U_IdUsuario int )  (int, error) {
	stmt := `INSERT INTO cartoes (idcartoes, Numero,Cvv, DataEmi,DataVal, Usuarios_idusuario) 
            VALUES(?,?,?,?,?,?)`

	result, err := m.DB.Exec(stmt, IdCartoes, Numero, Cvv, DataEmi, DataVal,U_IdUsuario)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *TravelModel) GetCartoes(id int) (*models.Cartoes, error) { //te, algum erro
	stmt := `SELECT idcartoes, Numero,Cvv, DataEmi,DataVal, Usuarios_idusuario FROM cartoes
           WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	s := &models.Cartoes{}

	err := row.Scan(&s.IdCartoes, &s.Numero, &s.Cvv, &s.DataEmi, &s.DataVal, &s.U_IdUsuario)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}
//volta depois

func (m *TravelModel) LatestCartoes() ([]*models.Cartoes, error) {
	stmt := `SELECT idcartoes, Numero,Cvv, DataEmi,DataVal, Usuarios_idusuario  FROM hoteis
          WHERE `

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	HoteisS := []*models.Hoteis{}
	for rows.Next() {
    s:= &models.Hoteis{}
    err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
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