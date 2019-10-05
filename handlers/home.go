package handlers

import (
	"net/http"
	"text/template"
)

// HomeHandler handles the Home Page of Dashboard 
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/home.html"))

	tmpl.Execute(w, nil)

}
