package handlers

import (
	"net/http"
	"text/template"

	kube "github.com/ErwinsExpertise/KubeDeploy/client"
)

//List deployments
func ListSiteHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/listsites.html"))
	var data Page

	dom := kube.GetDeployments()

	data.Sites = dom
	tmpl.Execute(w, data)
}
