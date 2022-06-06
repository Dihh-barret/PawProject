 package main

//go run cmd/web/*
import ( 
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

  Hoteis, err := app.Travels.MostPopularHoteis()
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
	err = ts.Execute(rw, Hoteis)
	if err != nil {
		app.serverError(rw, err)
		return
	}

}

//http://localhost:4000/snippet?id=1
func (app *application) showSnippet(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("idquartos"))
	if err != nil || id < 1 {
		app.notFound(rw)
		return
	}

  s, err := app.Travels.GetHoteis(id)
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
/*
func (app *application) createSnippet(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		rw.Header().Set("Allow", "POST")
		app.clientError(rw, http.StatusMethodNotAllowed)
		return
	}

	Nome := "NomedeHotel2"
	Cidade := "Cidade"
	Pais := "Pais2"
	expire := "7"
IdHoteis :=
  Nome string
  Cnpj 
  Cidade string
  Pais string
  S_IdSenhas int
  UpVotes int
  DownVotes int
  HospNum int 
	id, err := app.Travels.InsertHoteis()
	if err != nil {
		app.serverError(rw, err)
		return
	}

	http.Redirect(rw, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}*/