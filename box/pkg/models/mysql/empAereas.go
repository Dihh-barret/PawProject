package mysql

import (
	"database/sql"
	"merlin.com/box/pkg/models"
)

func (m *TravelModel) InsertEmp(IdEmpresasAerea int,
  Nome string, Cnpj string, Email string, Telefone string, Cidade string, Pais string)  (int, error) {
	stmt := `INSERT INTO empresasaereas (idEmpresasAereas, Nome, CNPJ, Telefone, Email, Cidade, Pais) 
            VALUES(?,?,?,?,?,?,?)`

	result, err := m.DB.Exec(stmt, IdEmpresasAerea, Nome, Cnpj, Email, Telefone, Cidade, Pais)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *TravelModel) GetEmp(id int) (*models.EmpresasAereas, error) { //te, algum erro
	stmt := `SELECT idEmpresasAereas, Nome, CNPJ, Telefone, Email, Cidade, Pais FROM itenscompras
           WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	s := &models.EmpresasAereas{}

	err := row.Scan(&s.IdEmpresasAerea, &s.Nome, &s.Cnpj, &s.Email, &s.Telefone, &s.Cidade, &s.Pais)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}
//volta depois

func (m *TravelModel) LatestEmp() ([]*models.EmpresasAereas, error) {
	stmt := `SELECT idEmpresasAereas, Nome, CNPJ, Telefone, Email, Cidade, Pais FROM itenscompras
           WHERE `

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	EmpresasAereasS := []*models.EmpresasAereas{}
	for rows.Next() {
    s:= &models.EmpresasAereas{}
    err = rows.Scan(&s.IdEmpresasAerea, &s.Nome, &s.Cnpj, &s.Email, &s.Telefone, &s.Cidade, &s.Pais)
    if err != nil{
      return nil, err
    }
    EmpresasAereasS = append(EmpresasAereasS, s)
	
  }
  err = rows.Err()
  if err != nil{ //executa funcao, coloca na variavel 
    return nil, err //executa
  }
	return nil, nil
}