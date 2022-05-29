//pra iniciar   go run cmd/web/*
package main

import (
	"database/sql"
	"flag"
	"log"
	"merlin.com/box/pkg/models/mysql"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infolog  *log.Logger
  Travel *mysql.TravelModel
}

//curl -i -X GET http://localhost:4000/snippet/create
func main() {
	//nome da flag, valor padrao e descricao
	addr := flag.String("addr", ":4000", "Porta da Rede")
	//concectar banco de dados
	dsn := flag.String("dsn","kbb8zd8UYz:LGKrbEczMK@tcp(remotemysql.com)/kbb8zd8UYz?parseTime=true","MySql Dsn") 

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERRO:\t",
		log.Ldate|log.Ltime|log.Lshortfile)
	//conectar banco de dados
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close() // se der ruim vai fechar o db

	//como se fosse objeto
	app := &application{
		errorLog: errorLog,
		infolog:  infoLog,
    Travel: &mysql.TravelModel{DB:db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Inicializando o servidor na porta %s\n", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
