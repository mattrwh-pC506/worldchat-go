package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var homeAddr = flag.String("home", ":8080", "chat landing page")

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("%v method not allowed", r.Method), http.StatusMethodNotAllowed)
	}
	http.ServeFile(w, r, "index.html")
}

func main() {
	password := flag.String("password", "", "Password for your chatroom.")
	flag.Parse()
	room := newRoom()
	go room.run()

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		authHandler(password, w, r)
	})
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(room, w, r)
	})
	err := http.ListenAndServe(*homeAddr, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
