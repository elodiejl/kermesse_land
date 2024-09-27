package services

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateTicketNumber() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000)) // Génère un numéro de ticket à 6 chiffres
}

func GetCurrentTime() string {
	return time.Now().Format("2024-01-02 15:04:05")
}
