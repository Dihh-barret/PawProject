package main

// curl -i -X POST http://localhost:4000/snippet/create
import (
	"fmt"
	"html/template" 
	"net/http"
	"strconv"
)


func(app *App) home(w http.ResponseWriter, r *http.Request){
  if r.URL.Path!= "/"{
    app.notFound(w)
    return
  }

  files:=[]string{
  "./ui/html/home.page.tmpl.html",
  "./ui/html/base.layout.tmpl.html",
  "./ui/html/foder.partial.tmpl.html",
  }
  ts , err := template.ParseFiles(files...)
  if err !=nil{
    app.errorLog.Println(err.Error())
    http.Error(w,"Internal Error",500)
    return
  }
  err = ts.Execute(w, nil)
  if err !=nil{
    app.errorLog.Println(err.Error())
    app.serverError(w, err)
    return
  }
}
//http://localhost:4000/snippet?=id123
func(app *App) ShowSnippet (w http.ResponseWriter, r *http.Request){
  id,err := strconv.Atoi(r.URL.Query().Get("id"))
  if err!= nil|| id<1{
    app.notFound(w)
    return
    }
  fmt.Fprintf(w,"Exibir Snippet ID: %d", id)
  }

func(app *App) CreateSnippet (w http.ResponseWriter, r *http.Request){
  
  if r.Method!= "POST"{
    w.Header().Set("Allow", "POST")
    app.clientError(w, http.StatusMethodNotAllowed)
    return
  }
  
  w.Write([]byte("Criar snippet"))

}
