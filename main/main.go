package main

//https://github.com/s1s1ty/go-mysql-crud
import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mintutu/restful-golang/driver"
	ph "github.com/mintutu/restful-golang/handler/http"
)

func main() {
	dbName := "db"
	dbPass := "123456"
	dbHost := "localhost"
	dbPort := "3306"

	connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	pHandler := ph.NewPostHandler(connection)
	cHandler := ph.NewCommentHandler()
	r.Get("/comments", cHandler.Fetch)
	r.Get("/posts", pHandler.Fetch)
	r.Post("/posts", pHandler.Create)
	r.Mount("/debug", middleware.Profiler())

	fmt.Println("Server listen at :8005")
	http.ListenAndServe(":8005", r)
}
