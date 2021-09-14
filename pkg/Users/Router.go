package Users

import (
	"github.com/go-chi/chi"
)

func UsersRoutes() *chi.Mux {
	router := chi.NewRouter() 
	router.Get("/getall", getHandler)
	router.Post("/register", postHandler)
	return router
}