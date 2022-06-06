package main

import "net/http"

func (app *application) routes() http.Handler {
  
  	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)  
	mux.HandleFunc("/Quarto", app.showSnippet)
	//mux.HandleFunc("/Hoteis/create", app.createSnippet)
  
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
 return app.logResquest(secureHeaders(mux))
}