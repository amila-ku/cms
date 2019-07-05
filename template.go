package cms

import (
	"html/template"
	"time"
)

//Tmpl is a exported variable with parsed template
var Tmpl = template.Must(template.ParseGlob("../templates/*"))

//Page struct defines the structure of a page
type Page struct {
	ID      int
	Title   string
	Content string
	Posts   []Post
}

//Post defines structure of post
type Post struct {
	ID            int
	Title         string
	Content       string
	DatePublished time.Time
	Comments      []Comment
}

//Comment defines structure of comment
type Comment struct {
	ID      int
	Message string
	Author  string
}
