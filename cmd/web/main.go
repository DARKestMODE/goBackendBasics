package main

import (
	"flag"
	"github.com/golangcollege/sessions"
	"github.com/jackc/pgx"
	"html/template"
	"log"
	"module1/pkg/models/postgre"
	"net/http"
	"os"
	"time"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      *postgre.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	config := pgx.ConnConfig{
		Host:     "db",
		Port:     5432,
		Database: "postgres",
		User:     "postgres",
		Password: "postgres",
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

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		snippets:      &postgre.SnippetModel{Conn: conn},
		templateCache: templateCache,
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on port %s", *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}
