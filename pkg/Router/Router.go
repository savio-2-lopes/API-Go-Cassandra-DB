package Router

import (
	database "home/lopes/Downloads/apache-cassandra-4.0.1/cassandra_project/backend/pkg/Database"
	"home/lopes/Downloads/apache-cassandra-4.0.1/cassandra_project/backend/pkg/Users"
	"fmt"
	"net/http"
	"log"
	"github.com/go-chi/chi"
)

func StartServer() *chi.Mux {
	database.SetupDBConnection()

	router := chi.NewRouter()
	router.Mount("/api/users", Users.UsersRoutes())
	fmt.Println("Servidor rodando na porta 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", router))
	return router
}