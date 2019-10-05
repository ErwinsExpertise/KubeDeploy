package handlers

import (
	"net/http"
	"text/template"

	kube "github.com/ErwinsExpertise/KubeDeploy/client"
)

//NewSiteHandler Handles Creation of new sites
func NewSiteHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/addsite.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	var data Page

	domain := r.FormValue("domain")
	password := r.FormValue("password")

	dom := kube.CreateSite(domain, password)
	data.Message = dom
	data.Success = true

	tmpl.Execute(w, data)
}

type Page struct {
	Title   string
	Message string
	Sites   []string
	Success bool
}
