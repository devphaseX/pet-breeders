package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
}

type appConfig struct {
	useCache bool
}

func main() {
	app := application{}

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}

	fmt.Println("server running on port", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
