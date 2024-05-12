package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adeindra6/intikom_test/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)
	http.Handle("/", r)

	localServer := "http://localhost:8080"
	fmt.Println(fmt.Sprintf("Server berjalan pada %s", localServer))
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
