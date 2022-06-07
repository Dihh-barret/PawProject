package entity

//Person object for REST(CRUD)
type Person struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Idade     int    `json:"idade"`
	Profissao string `json:"profissao"`
}
