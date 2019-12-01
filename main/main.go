package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mintutu/restful-golang/containers"
)

func main() {
	container := containers.BuildContainer()
	err := container.Invoke(func(router *chi.Mux) {
		fmt.Println("Server listen at :8005")
		http.ListenAndServe(":8005", router)
	})

	if err != nil {
		panic(err)
	}
}
