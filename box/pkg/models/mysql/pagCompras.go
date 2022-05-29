package mysql

import (
	"database/sql"
	"merlin.com/box/pkg/models"
  "time"
)
 
func (m *TravelModel) InsertPagCom( IdPagamentosCompras int , DataPag time.Timer, ValorTotal float32 , U_IdUsuario int, TipoPag_IdTipoPagamento int )  (int, error) {
	stmt := `INSERT INTO pagamentos_compras (idPagamentos_Compras, ValorTotal, DataPagamento, Usuarios_idusuario, Tipo_pagamentos_idTipo_pagamentos) 
            VALUES(?,?,?,?,?)`

	result, err := m.DB.Exec(stmt, IdPagamentosCompras, ValorTotal, DataPag,  U_IdUsuario , TipoPag_IdTipoPagamento)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *TravelModel) GetPagCom(id int) (*models.PagamentosCompras, error) { //te, algum erro
	stmt := `SELECT idPagamentos_Compras, ValorTotal, DataPagamento, Usuarios_idusuario, Tipo_pagamentos_idTipo_pagamentos FROM pagamentos_compras
           WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	s := &models.PagamentosCompras{}

	err := row.Scan(&s.IdPagamentosCompras, &s.ValorTotal, &s.DataPag, &s.U_IdUsuario, &s.TipoPag_IdTipoPagamento)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}
//volta depois

func (m *TravelModel) LatestPagCom() ([]*models.PagamentosCompras, error) {
	stmt := `SELECT idPagamentos_Compras, ValorTotal, DataPagamento, Usuarios_idusuario, Tipo_pagamentos_idTipo_pagamentos FROM pagamentos_compras
          WHERE `

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	PagamentosComprasS := []*models.PagamentosCompras{}
	for rows.Next() {
    s:= &models.PagamentosCompras{}
    err = rows.Scan(&s.IdPagamentosCompras, &s.ValorTotal, &s.DataPag, &s.U_IdUsuario, &s.TipoPag_IdTipoPagamento)
    if err != nil{
      return nil, err
    }
    PagamentosComprasS = append(PagamentosComprasS, s)
	
  }
  err = rows.Err()
  if err != nil{ //executa funcao, coloca na variavel 
    return nil, err //executa
  }
	return nil, nil
}