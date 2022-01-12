package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shreyasir/golang/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	pdt := handlers.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/products", pdt)

	s := &http.Server{
		Addr:         ":8091",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	l.Fatal(s.ListenAndServe())
}
