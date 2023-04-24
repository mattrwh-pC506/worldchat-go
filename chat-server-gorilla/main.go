package main

import (
	"crypto/sha256"
	"flag"
	"log"
	"net/http"
	"os"

	"chat-server-gorilla/handlers"
)

func main() {
	homeAddr := flag.String("home", ":8080", "chat landing page")
	password := os.Getenv("WCG_PASSWORD")

	// Weird I know, just a hack to generate a secret key given the weirdness of passing a password as a flag
	hash := sha256.Sum256([]byte(password))
	secret := hash[:32]
	flag.Parse()

	log.Printf("Running chat-server-gorilla on %v\n", *homeAddr)

	room := handlers.NewRoom()
	go room.Run()

	mux := http.NewServeMux()

	// Static Assets
	fileServer := http.FileServer(http.Dir("../chat-ui/build"))
	mux.Handle("/static/", fileServer)
	mux.Handle("/manifest.json", fileServer)
	mux.Handle("/robots.txt", fileServer)
	mux.Handle("/favicon.ico", fileServer)
	mux.Handle("/logo192.png", fileServer)
	mux.Handle("/logo512.png", fileServer)

	// index.html
	mux.HandleFunc("/", handlers.IndexHandler)

	// REST
	mux.HandleFunc("/login", handlers.LoginHandler(password, secret))
	mux.HandleFunc("/auth", handlers.AuthHandler(secret))

	// Websockets
	mux.HandleFunc("/ws", handlers.ChatHandler(room))

	// Serve on Home Address
	err := http.ListenAndServe(*homeAddr, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
