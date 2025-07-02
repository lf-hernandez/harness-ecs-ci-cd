package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Printf("Request received for path: %s\n", path)
	
	if path == "/" {
		fmt.Fprintf(w, "Hello! You're at the root endpoint. Version: %s\n", getVersion())
	} else {
		fmt.Fprintf(w, "You requested path: %s (Version: %s)\n", path, getVersion())
	}
}

func getVersion() string {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		return "dev"
	}
	return version
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)
	
	fmt.Printf("Server starting on port %s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
