package main
// curl -i -X POST http://localhost:4000/snippet/create
import ( 
	"net/http" 
)


func (app * App) routes() *http.ServeMux{
  
	mux:= http.NewServeMux()
  mux.HandleFunc("/", app.home)
  mux.HandleFunc("/snippet", app.ShowSnippet)
  mux.HandleFunc("/snippet/create", app.CreateSnippet)
  fileServer:=http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/",http.StripPrefix("/static", fileServer))
   return mux
}