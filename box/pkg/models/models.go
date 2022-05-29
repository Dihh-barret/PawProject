//colocar as coisas do banco de dados aqui trazendo para o programa e trazendo para uma variavel

package models

import ("time"
       "errors")

var ErrNoRecord = errors.New("models: no matching record Found")

type Usuarios struct{
  Idusuario int
  Nome string
  Email string
  Telefone string
  Cpf int
       //estrutura que o html recebeyu
}
type Comentarios struct{
  IdComentarios int
  Texto string
  DataComent time.Timer
  Upvote int
  Downvote int
  U_IdUsuario int
  H_IdHotel int 
}
type Quartos struct{
  IdQuartos int
  Numero int 
  Upvote int
  Downvote int
  Preço float32
  Status int
  H_IdHotel int 
}
type ItensCompras struct{
  IdItens int
  PrecoTotal float32 
  Pag_IdPagamentos int
  Q_H_IdHoteis int
  Q_IdQuartos int
  Pass_IdPassagem int 
  Pass_Emp_IdEmpresasAerea int 
}
type PagamentosCompras struct{
  IdPagamentosCompras int 
  DataPag time.Timer
  ValorTotal float32 
  U_IdUsuario int
  TipoPag_IdTipoPagamento int 
}
type TipoPagamento struct{
  IdTipoPagamento int
  Tipo string 
}
type Cartoes struct{
  IdCartoes int
  Numero string
  DataVal time.Timer
  DataEmi time.Timer
  Cvv int 
  U_IdUsuario int 
}
type Hoteis struct{
  IdHoteis int
  Nome string
  Cidade string
  Pais string
  Cnpj string
}
type Fotos struct{
  IdFotos int
  Imagem string  
  DataEnvio time.Timer
  H_IdHoteis int
  Q_IdQuartos int
}
type Passagens struct{
  IdPassagem int
  Numero int 
  DataEmbarque time.Timer 
  Preço float32
  Status int
  Emp_IdEmpresasAerea int 
}
type EmpresasAereas struct{
  IdEmpresasAerea int
  Nome string
  Email string
  Telefone string
  Cidade string
  Pais string
  Cnpj string
}