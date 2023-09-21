package main

import (
	"laptop/controller"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/students", controller.All)
	r.Post("/students", controller.Create)
	r.Put("/students/{Sid}", controller.Update)
	r.Delete("/students/{Sid}", controller.Delete)
	http.ListenAndServe(":4000", r)
}
