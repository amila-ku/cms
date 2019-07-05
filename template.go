package cms

import (
	"html/template"
	"time"
)

//Tmpl is a exported variable with parsed template
var Tmpl = template.Must(template.ParseGlob("../templates/*"))

//Page struct defines the structure of a page
type Page struct {
	Title   string
	Content string
	Posts   []Post
}

type Post struct {
	Title         string
	Content       string
	DatePublished time.Time
	Comments      []Comment
}

type Comment struct {
	Message string
	Author  string
}
