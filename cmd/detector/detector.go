package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var MessageCh = make(chan []byte)

func main() {
	go runWebSocketListener()

	ReadMessage("localhost:9094")
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Allow all connections for demonstration purposes.
			return true
		},
	}
)

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()
	for {
		select {
		case message := <-MessageCh:
			err = conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Failed to send message:", err)
				return
			}

			log.Println("Message sent:", message)
			break
		}
	}
}

func runWebSocketListener() {
	http.HandleFunc("/ws", websocketHandler)

	// Start the HTTP server.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
