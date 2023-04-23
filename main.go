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
	http.ServeFile(w, r, "./chat-ui/build/index.html")
}

func main() {
	password := flag.String("password", "", "Password for your chatroom.")
	flag.Parse()
	room := newRoom()
	go room.run()

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	fileServer := http.FileServer(http.Dir("./chat-ui/build"))
	mux.Handle("/static/", fileServer)
	mux.Handle("/manifest.json", fileServer)
	mux.Handle("/robots.txt", fileServer)
	mux.Handle("/favicon.ico", fileServer)
	mux.Handle("/logo192.png", fileServer)
	mux.Handle("/logo512.png", fileServer)

	mux.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		authHandler(password, w, r)
	})

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chatHandler(room, w, r)
	})
	err := http.ListenAndServe(*homeAddr, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
