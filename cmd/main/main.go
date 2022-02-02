package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/usamah/go-workout/pkg/routes"
	_ "gorm.io/gorm"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Alons y")
	routes.RegisterWorkoutRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3007", r))

}
