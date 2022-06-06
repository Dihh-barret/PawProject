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

func (m *TravelModel) CartoesUsurio(id int) ([]*models.Cartoes, error) {
	stmt := `SELECT * FROM cartoes INNER JOIN usuarios on cartoes.usuarios.idsuarios = usuarios.idsuarios 
          WHERE usuarios.idsuarios = ?`

	rows, err := m.DB.Query(stmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	CartoesS := []*models.Cartoes{}
	for rows.Next() {
    s:= &models.Cartoes{}
    err = rows.Scan(&s.IdCartoes, &s.Numero, &s.Cvv, &s.DataEmi, &s.DataVal)
    if err != nil{
      return nil, err
    }
    CartoesS = append(CartoesS, s)
	
  }
  err = rows.Err()
  if err != nil{ //executa funcao, coloca na variavel 
    return nil, err //executa
  } 
	return nil, nil
}