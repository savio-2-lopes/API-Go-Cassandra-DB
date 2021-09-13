package Users

import (
	database "home/lopes/Downloads/apache-cassandra-4.0.1/cassandra_project/backend/pkg/Database"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	user User
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Você tem o user1")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&user)
	addUser(&user)

	fmt.Println("Adicionado o usuário ao banco de dados ...")
	json.NewEncoder(w).Encode("Você adicionou um novo usuario para a Cassandra DB")
}

func addUser(user *User) {
	query := `INSERT INTO users(first_name, last_name, email) VALUES (?,?,?)`
	database.ExecuteQuery(query, user.FirstName, user.LastName, user.Email)
}