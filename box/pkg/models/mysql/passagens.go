package mysql

import (
	"database/sql"
	"time"

	"merlin.com/box/pkg/models"
)

 
func (m *TravelModel) InsertPass(IdPassagem int, Numero int, DataEmbarque time.Timer, Preço float32, Status int,  Emp_IdEmpresasAerea int )  (int, error) {
	stmt := `INSERT INTO passagens (idpassagens, NumeroDaPassagem, DataEmbarque, Preco, EmpresasAereas_idEmpresasAereas) 
            VALUES(?,?,?,?,?)`

	result, err := m.DB.Exec(stmt, IdPassagem, Numero, DataEmbarque, Preço, Status,Emp_IdEmpresasAerea)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *TravelModel) GetPass(id int) (*models.Passagens, error) { //te, algum erro
	stmt := `SELECT idpassagens, NumeroDaPassagem, DataEmbarque, Preco, EmpresasAereas_idEmpresasAereas FROM passagens
           WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	s := &models.Passagens{}

	err := row.Scan(&s.IdPassagem, &s.Numero, &s.DataEmbarque, &s.Preço, &s.Status, &s.Emp_IdEmpresasAerea)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}


