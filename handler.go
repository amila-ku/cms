package cms

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

//HelloServer servers hello
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

//ServeIndex serves index page
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	page := Page{
		Title:   "Hello Web",
		Content: "Web Page with go templates being shown",
		Posts: []Post{
			Post{
				Title:         "First Post",
				Content:       "My First Post",
				DatePublished: time.Now(),
			},
			Post{
				Title:         "Second Post",
				Content:       "My Second Post",
				DatePublished: time.Now().Add(-time.Hour),
				Comments: []Comment{
					Comment{
						Message: "This is great",
						Author:  "maxyyy",
					},
					Comment{
						Message: "Post more",
						Author:  "Randy",
					},
				},
			},
		},
	}

	Tmpl.ExecuteTemplate(w, "index", page)

}

//HandleNew handles new page creation
func HandleNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Tmpl.ExecuteTemplate(w, "new", nil)
	case "POST":
		title := r.FormValue("title")
		content := r.FormValue("content")
		contentType := r.FormValue("content-type")
		r.ParseForm()

		if contentType == "post" {
			Tmpl.ExecuteTemplate(w, "post", Post{
				Title:   title,
				Content: content,
			})
			return
		}
		if contentType == "page" {
			page := Page{
				Title:   title,
				Content: content,
			}
			Tmpl.ExecuteTemplate(w, "page", page)

			_, err := CreatePage(page)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

	default:
		http.Error(w, "Method Not Supported"+r.Method, http.StatusMethodNotAllowed)

	}

}

//HandlePost handles new page creation
func HandlePost(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "post")

	if path == "" {
		http.NotFound(w, r)
		return
	}

	post := Post{
		Title:   "New Post",
		Content: "Just another post nw",
	}

	Tmpl.ExecuteTemplate(w, "post", post)

}

//HandlePage handles new page creation
func HandlePage(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "page")

	if path == "" {
		pages, err := GetPage()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		Tmpl.ExecuteTemplate(w, "page", pages)
		return
	}

	page := Page{
		Title:   "New Page with contetn",
		Content: "Just another post nw post from Dortmund",
	}

	Tmpl.ExecuteTemplate(w, "page", page)

}
