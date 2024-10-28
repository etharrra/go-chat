package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"

	"github.com/gofiber/websocket/v2"
)

type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	broadcast chan *Message
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *Message),
	}
}

/**
 * HandleWebSocket handles incoming messages from a WebSocket connection.
 * It registers a new client, reads incoming messages,
 * unmarshals them into a Message struct,
 * and broadcasts the message to all clients.
 * @param c The WebSocket connection to handle
 */
func (s *WebSocketServer) HandleWebSocket(c *websocket.Conn) {
	// register new client
	s.clients[c] = true
	defer func() {
		delete(s.clients, c)
		c.Close()
	}()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("Read Error: ", err)
			break
		}

		// send the message to broadcast channels
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Fatalf("Error Unmarshaling")
		}
		s.broadcast <- &message
	}
}

/**
 * HandleMessage sends a message to all connected clients
 * using the WebSocketServer's broadcast channel.
 * If an error occurs while sending the message to a client,
 * the client is closed and removed from the server's client list.
 */
func (s *WebSocketServer) HandelMessage() {
	for {
		msg := <-s.broadcast

		// send message to all clients
		for client := range s.clients {
			err := client.WriteMessage(websocket.TextMessage, getMessageTemplate(msg))
			if err != nil {
				log.Printf("Write Error: %v", err)
				client.Close()
				delete(s.clients, client)
			}
		}
	}
}

/**
 * getMessageTemplate generates a byte slice by rendering the message template with the provided message data.
 * 
 * @param msg: Pointer to the Message struct containing data to be rendered in the template.
 * @return []byte: Byte slice representing the rendered message template.
 */
func getMessageTemplate(msg *Message) []byte {
	tmpl, err := template.ParseFiles("views/message.html")
	if err != nil {
		log.Fatalf("template parsing: %s", err)
	}

	// render the template with message as data
	var renderMessage bytes.Buffer
	err = tmpl.Execute(&renderMessage, msg)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	return renderMessage.Bytes()
}
