package forms

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.Method != "Post" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fname := r.FormValue("firster")
	lname := r.FormValue("laster")

	d := struct {
		First string
		Last  string
	}{
		First: fname,
		Last:  lname,
	}

	tpl.ExecuteTemplate(w, "process.html", d)
}
