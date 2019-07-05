package cms

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "cms"
)

//Pgstore is a type to initiate a postgres db connection
type Pgstore struct {
	DB *sql.DB
}

func newDB() *Pgstore {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// fmt.Println("Successfully connected!")

	return &Pgstore{
		DB: db,
	}
}

//CreatePage saves page to db
func CreatePage(p Page) (int, error) {
	var id int
	err := newDB().DB.QueryRow("INSERT INTO PAGES(title,content) values($1,$2) RETURNING id", p.Title, p.Content).Scan(&id)
	return id, err
}

//GetPage returns pages
func GetPage() ([]Page, error) {
	rows, err := newDB().DB.Query("SELECT * FROM PAGES")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	pages := []Page{}

	for rows.Next() {
		var p Page
		err := rows.Scan(p.ID, p.Title, p.Content)
		if err != nil {
			return nil, err
		}
		pages = append(pages, p)
	}

	return pages, nil
}
