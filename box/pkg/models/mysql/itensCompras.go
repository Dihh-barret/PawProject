package mysql

import (
	"database/sql"
	"merlin.com/box/pkg/models"
)
 
func (m *TravelModel) InsertItens(IdItens int, PrecoTotal float32 , Pag_IdPagamentos int, Q_H_IdHoteis int, Q_IdQuartos int, Pass_IdPassagem int, Pass_Emp_IdEmpresasAerea int )  (int, error) {
	stmt := `INSERT INTO itenscompras ( idItensCompras, PrecoTotal, Pagamentos_Compras_idPagamentos_Compras, Quartos_idquartos, Quartos_Hoteis_idhoteis) 
            VALUES(?,?,?,?,?)`

	result, err := m.DB.Exec(stmt, IdItens, PrecoTotal, Pag_IdPagamentos, Q_IdQuartos, Q_H_IdHoteis)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *TravelModel) GetItens(id int) (*models.ItensCompras, error) { //te, algum erro
	stmt := `SELECT * FROM itenscompras
           WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	s := &models.ItensCompras{}

	err := row.Scan(&s.IdItens, &s.PrecoTotal, &s.Pag_IdPagamentos, &s.Q_IdQuartos, &s.Q_H_IdHoteis)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}
//volta depois

func (m *TravelModel) LatestItens() ([]*models.ItensCompras, error) {
	stmt := `SELECT * FROM itenscompras
           WHERE `

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	ItensComprasS := []*models.ItensCompras{}
	for rows.Next() {
    s:= &models.ItensCompras{}
    err = rows.Scan(&s.IdItens, &s.PrecoTotal, &s.Pag_IdPagamentos, &s.Q_IdQuartos, &s.Q_H_IdHoteis)
    if err != nil{
      return nil, err
    }
    ItensComprasS = append(ItensComprasS, s)
	
  }
  err = rows.Err()
  if err != nil{ //executa funcao, coloca na variavel 
    return nil, err //executa
  }
	return nil, nil
}