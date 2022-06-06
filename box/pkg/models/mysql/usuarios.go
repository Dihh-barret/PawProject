package mysql

import (
	"database/sql"
	"merlin.com/box/pkg/models"
)

  

func (m *TravelModel) InsertUser(idusuario int, Nome string, Email string, Cpf string,Telefone string)  (int, error) {
	stmt := `INSERT INTO usuarios (idusuario, Nome, Email, Cpf, Telefone) 
            VALUES(?,?,?,?,?)`

	result, err := m.DB.Exec(stmt, idusuario, Nome, Email, Cpf, Telefone)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *TravelModel) GetUser(id int) (*models.Usuarios, error) { //te, algum erro
	stmt := `SELECT idusuario, Nome, Email, Cpf, Telefone FROM usuarios
           WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	s := &models.Usuarios{}

	err := row.Scan(&s.Idusuario, &s.Nome, &s.Email, &s.Cpf, &s.Telefone)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}

func (m *TravelModel) UsuarioPagCom(id int) ([]*models.PagamentosCompras, error) {
	stmt := `SELECT * FROM pagamentos_compras INNER JOIN idusuario ON pagamentos_compras.Usuarios_idusuario = usuarios.idusuario
          WHERE usuarios.idusuario = ?`

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


