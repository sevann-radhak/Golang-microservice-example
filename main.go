package main

import (
	// "handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sevannr/Golang-microservice-example/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	fmt.Println(l)
	hh := handlers.NewHello(l) //handlers.NewHello(l)
	// gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	// sm.Handle("/", gh)

	http.ListenAndServe(":9090", sm)
}
