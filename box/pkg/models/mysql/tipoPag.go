package mysql

import (
	"database/sql"
	"merlin.com/box/pkg/models"
  "time"
) 
func (m *TravelModel) InsertTipo(IdTipoPagamento int, Tipo string )  (int, error) {
	stmt := `INSERT INTO idTipo_pagamentos, Tipo) 
            VALUES(?,?)`

	result, err := m.DB.Exec(stmt, IdTipoPagamento, Tipo)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *TravelModel) GetTipo(id int) (*models.TipoPagamento, error) { //te, algum erro
	stmt := `SELECT dTipo_pagamentos, Tipo FROM tipo_pagamentos
           WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	s := &models.TipoPagamento{}

	err := row.Scan(&s.IdTipoPagamento, &s.Tipo)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}
//volta depois

func (m *TravelModel) LatestTipo() ([]*models.TipoPagamento, error) {
	stmt := `SELECT dTipo_pagamentos, Tipo FROM tipo_pagamentos
           WHERE  `

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	TipoPagamentoS := []*models.TipoPagamento{}
	for rows.Next() {
    s:= &models.TipoPagamento{}
    err = rows.Scan(&s.IdTipoPagamento, &s.Tipo)
    if err != nil{
      return nil, err
    }
    TipoPagamentoS = append(TipoPagamentoS, s)
	
  }
  err = rows.Err()
  if err != nil{ //executa funcao, coloca na variavel 
    return nil, err //executa
  }
	return nil, nil
}