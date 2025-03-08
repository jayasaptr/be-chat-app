package handlers

import (
	"chat-app/config"
	"chat-app/models"
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
)

var clients = make(map[*websocket.Conn]string) // Simpan user_id untuk setiap koneksi
var broadcast = make(chan models.Message)      // Channel untuk pesan masuk

// Inisialisasi WebSocket Broadcast
func InitWebSocket() {
	go func() {
		for {
			msg := <-broadcast
			fmt.Printf("Broadcasting message: %+v\n", msg)

			for client, userID := range clients {
				if userID == msg.ReceiverID || userID == msg.SenderID {
					err := client.WriteJSON(msg)
					if err != nil {
						log.Println("Error sending message:", err)
						client.Close()
						delete(clients, client)
					}
				}
			}
		}
	}()
}

// WebSocket Chat Handler dengan autentikasi JWT
func ChatHandler(c *websocket.Conn) {
	userID, ok := c.Locals("user_id").(string)
	if !ok || userID == "" {
		log.Println("❌ WebSocket ditutup: user_id tidak ditemukan")
		c.Close()
		return
	}

	log.Println("✅ WebSocket terhubung untuk user:", userID)

	clients[c] = userID
	defer func() {
		delete(clients, c)
		log.Println("❌ WebSocket ditutup untuk user:", userID)
		c.Close()
	}()

	for {
		var msg models.Message
		err := c.ReadJSON(&msg)
		if err != nil {
			log.Println("❌ Error membaca pesan:", err)
			break
		}

		msg.SenderID = userID

		// Simpan pesan ke database
		db := config.GetDB()
		if err := db.Create(&msg).Error; err != nil {
			log.Println("❌ Error menyimpan pesan ke DB:", err)
		}

		// Kirim pesan ke semua client
		broadcast <- msg
	}
}
