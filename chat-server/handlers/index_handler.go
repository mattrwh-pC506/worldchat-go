package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("URL: %v", r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("%v method not allowed", r.Method), http.StatusMethodNotAllowed)
	}

	dir, _ := os.Getwd()
	log.Printf("DIR", dir)
	http.ServeFile(w, r, "../chat-ui/build/index.html")
}
