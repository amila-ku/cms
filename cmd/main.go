package main

import (
	"net/http"

	"github.com/cms"
)

func main() {
	// page := cms.Page{
	// 	Title:   "Hello Web",
	// 	Content: "Web Page with go templates",
	// }

	// cms.Tmpl.ExecuteTemplate(os.Stdout, "index", page)
	http.HandleFunc("/", cms.HandleIndex)
	http.HandleFunc("/new", cms.HandleNew)
	http.HandleFunc("/page", cms.HandlePage)
	http.HandleFunc("/post", cms.HandlePost)
	http.ListenAndServe(":9000", nil)

}
