package main

import (
	"home/lopes/Downloads/apache-cassandra-4.0.1/cassandra_project/backend/pkg/Router"
	"fmt"
)

func main() {
	fmt.Println("Iniciando servidor ...")
	Router.StartServer()
}