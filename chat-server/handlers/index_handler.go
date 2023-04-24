package handlers

import (
	"fmt"
	"log"
	"net/http"
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

	http.ServeFile(w, r, "../chat-ui/build/index.html")
}
