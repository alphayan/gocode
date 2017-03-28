package main

// import (
// 	_ "fmt"
// 	"github.com/gorilla/websocket"
// 	_ "log"
// 	"net/http"
// )

// var clients = make(map[*websocket.Conn]bool) // connected clients
// var broadcast = make(chan Message)           // broadcast channel
// var upgrader = websocket.Upgrader{}

// type Message struct {
// 	Email    string `json:"email"`
// 	Username string `json:"username"`
// 	Message  string `json:"message"`
// }

// func Websocket() {
// 	// Create a simple file server
// 	fs := http.FileServer(http.Dir("../public"))
// 	http.Handle("/", fs)
// 	// Configure websocket route
// 	http.HandleFunc("/ws", handleConnections)
// 	// Start listening for incoming chat messages
// 	go handleMessages()
// }
