package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env load wrong")
		return
	}
	port := os.Getenv("Port")
	mux := http.NewServeMux()
	mux.HandleFunc("/testAlive", TestAlive)
	mux.HandleFunc("/getNodes", GetNodes)
	server := &http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}
	log.Println("start server at:", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Println("server start error", err)
		return
	}
}
