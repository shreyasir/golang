package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	sm := http.ServeMux(hh)
	sm.Handle(hh)
	http.ListenAndServe(":8090", sm)
}
