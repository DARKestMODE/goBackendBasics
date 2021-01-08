package main

import (
	"flag"
	"github.com/jackc/pgx"
	"html/template"
	"log"
	"module1/pkg/postgre"
	"net/http"
	"os"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *postgre.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	config := pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "golang",
		User:     "golang_user",
		Password: "1golang1",
	}

	conn, err := pgx.Connect(config)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &postgre.SnippetModel{Conn: conn},
		templateCache: templateCache,
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}
