package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket Upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (be cautious in production)
	},
}

// WebSocket handler
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected!")

	// Read and echo messages
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		fmt.Printf("Received: %s\n", message)

		// Echo the message back to the client
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}

func main() {
	// Serve static files (HTML, etc.)
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	// Handle WebSocket connections
	http.HandleFunc("/ws", handleWebSocket)

	// Start the server
	fmt.Println("Server started at http://localhost:1080")
	log.Fatal(http.ListenAndServe(":1080", nil))
}
