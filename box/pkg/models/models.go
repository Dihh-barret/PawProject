//colocar as coisas do banco de dados aqui trazendo para o programa e trazendo para uma variavel

package models

import ("time"
       "errors")

var ErrNoRecord = errors.New("models: no matching record Found")

type Snippet struct{
  ID int
  Title string
  Content string
  Created time.Time
  Expires time.Time
  Test string
  
}
type Usuarios struct{
  Idusuario int
  Nome string
  Email string
  Telefone string
  Cpf int
  S_IdSenhas int
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
type DataReservas struct{
  IdDataReservas int 
  CheckIn time.Timer 
  CheckOut time.Timer 
  IdQuartos int
  Q_H_IdHotel int 
}
type Quartos struct{
  IdQuartos int
  Numero int 
  Preço float32 
  H_IdHotel int 
  Descricao string
  Wifi bool
  ArCond bool
}
type ItensCompras struct{
  IdItens int
  PrecoTotal float32 
  Pag_IdPagamentos int
  Q_H_IdHoteis int
  Q_IdQuartos int 
  Pass_IdPassagem int
}
type PagamentosCompras struct{
  IdPagamentosCompras int 
  ValorTotal float32 
  DataPag time.Timer
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
  Cnpj string 
  Cidade string
  Pais string
  S_IdSenhas int
  UpVotes int
  DownVotes int
  HospNum int 
  
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
  Emp_IdEmpresasAerea int  
  Itens_IdeItens int  
  
}
type EmpresasAereas struct{
  IdEmpresasAerea int
  Nome string
  Cnpj string
  Telefone string
  Email string
  Cidade string
  Pais string
  S_IdSenhas int
  
}
type Senhas struct{
  IdSenhas int
  Senha string 
  
}