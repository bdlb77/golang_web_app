package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // driver for Postgres
	"golang_web_app/src/github.com/bdlb77/webapp/controller"
	"golang_web_app/src/github.com/bdlb77/webapp/middleware"
	"golang_web_app/src/github.com/bdlb77/webapp/model"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// populating templates (view parsing and finding)
	templates := populateTemplates()
	// handles controller layer
	controller.StartUp(templates)
	// connect to DB
	db := connectToDB()
	// gracefully close DB when app shuts down
	defer db.Close()
	// serve port
	http.ListenAndServe(":8000", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})
}

func connectToDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://bryanleighton:@localhost/golang_db?sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("Could not connect to DB: %v", err))
	}
	model.SetDatabase(db)
	return db
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
