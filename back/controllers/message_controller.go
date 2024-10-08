package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader permet de mettre à niveau la connexion HTTP en connexion WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client représente un utilisateur connecté via WebSocket
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}

// MessageController gère les connexions WebSocket et la messagerie
type MessageController struct {
	Clients   map[string]*Client // Stocke les utilisateurs connectés
	Broadcast chan []byte        // Canal de diffusion des messages
}

// Créer un nouveau contrôleur de chat
func NewMessageController() *MessageController {
	return &MessageController{
		Clients:   make(map[string]*Client),
		Broadcast: make(chan []byte),
	}
}

// Gestionnaire pour les connexions WebSocket
func (ctrl *MessageController) WebSocketHandler(c *gin.Context) {
	// Upgrader la connexion HTTP en WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set WebSocket upgrade"})
		return
	}

	clientID := c.Query("client_id") // L'organisateur ou le teneur de stand doit fournir son ID
	client := &Client{
		ID:     clientID,
		Socket: conn,
		Send:   make(chan []byte),
	}

	// Enregistrer le client
	ctrl.RegisterClient(client)

	// Gérer les messages envoyés par le client
	go ctrl.HandleMessages(client)

	// Gérer les messages reçus par le client
	go ctrl.ReceiveMessages(client)
}

// Fonction pour gérer l'envoi des messages
func (ctrl *MessageController) HandleMessages(client *Client) {
	defer func() {
		client.Socket.Close()
	}()

	for {
		message := <-client.Send
		if err := client.Socket.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Println("Erreur lors de l'envoi du message:", err)
			return
		}
	}
}

// Fonction pour gérer la réception des messages
func (ctrl *MessageController) ReceiveMessages(client *Client) {
	defer func() {
		err := client.Socket.Close()
		if err != nil {
			return
		}
	}()

	for {
		_, message, err := client.Socket.ReadMessage()
		if err != nil {
			fmt.Println("Erreur lors de la réception du message:", err)
			return
		}

		// Diffuser le message aux autres utilisateurs
		ctrl.BroadcastMessage(client.ID, message)
	}
}

// Enregistrer un nouveau client
func (ctrl *MessageController) RegisterClient(client *Client) {
	ctrl.Clients[client.ID] = client
	fmt.Printf("Client %s connecté\n", client.ID)
}

// Diffuser un message à tous les clients connectés
func (ctrl *MessageController) BroadcastMessage(senderID string, message []byte) {
	for id, client := range ctrl.Clients {
		if id != senderID { // Ne pas envoyer le message à l'expéditeur
			client.Send <- message
		}
	}
}

// Supprimer un client lorsqu'il se déconnecte
func (ctrl *MessageController) RemoveClient(clientID string) {
	delete(ctrl.Clients, clientID)
	fmt.Printf("Client %s déconnecté\n", clientID)
}
