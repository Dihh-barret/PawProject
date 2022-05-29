package mysql

import (
	"database/sql"
	"merlin.com/box/pkg/models"
)

type HoteisModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(IdHoteis int, Nome string, Cidade string, Pais string,Cnpj string)  (int, error) {
	stmt := `INSERT INTO hoteis (idHoteis, Nome, Email, Cpf, Telefone) 
            VALUES(?,?,UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

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

func (m *UserModel) Get(id int) (*models.Usuarios, error) { //te, algum erro
	stmt := `SELECT idHoteis, Nome, Email, Cpf, Telefone FROM hoteis
           WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	s := &models.Usuarios{}

	err := row.Scan(&s.IdHoteis, &s.Nome, &s.Cidade, &s.Pais, &s.Cnpj)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}
func (m *UserModel) Latest() ([]*models.Usuarios, error) {
	stmt := `SELECT id, title,content, created, expires FROM snippets
          WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//crando slice
	UsuariosS := []*models.Usuarios{}
	for rows.Next() {
    s:= &models.Usuarios{}
    err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
    if err != nil{
      return nil, err
    }
    UsuariosS = append(UsuariosS, s)
	
  }
  err = rows.Err()
  if err != nil{ //executa fincao, coloca na variavel 
    return nil, err //executa
  }
	return nil, nil
}