package main

//go run cmd/web/*
import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"merlin.com/box/pkg/models"
)

func (app *application) home(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(rw)
		return
	}

  snippets, err := app.snippets.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

	files := []string{
		"./ui/html/home.page.tmpl.html",
		"./ui/html/base.layout.tmpl.html",
		"./ui/html/footer.partial.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(rw, err)
		return
	}
	err = ts.Execute(rw, snippets)
	if err != nil {
		app.serverError(rw, err)
		return
	}

}

//http://localhost:4000/snippet?id=1
func (app *application) showSnippet(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(rw)
		return
	}

  s, err := app.snippets.Get(id)
  if err == models.ErrNoRecord {
    app.notFound(rw)
    return
  }else if err != nil{
    app.serverError(rw, err)
    return
  }
  
  files := []string{
		"./ui/html/show.page.tmpl.html",
		"./ui/html/base.layout.tmpl.html",
		"./ui/html/footer.partial.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(rw, err)
		return
	}
	err = ts.Execute(rw, s)
	if err != nil {
		app.serverError(rw, err)
		return
	}
  
}

func (app *application) createSnippet(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		rw.Header().Set("Allow", "POST")
		app.clientError(rw, http.StatusMethodNotAllowed)
		return
	}

	title := "Aula de hoje"
	content := "Tentando lidar com o banco de dados"
	expire := "7"

	id, err := app.snippets.Insert(title, content, expire)
	if err != nil {
		app.serverError(rw, err)
		return
	}

	http.Redirect(rw, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}