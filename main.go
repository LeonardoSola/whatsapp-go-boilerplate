package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/LeonardoSola/whatsapp-go-boilerplate/adapters/cmd/whastapp"
	handlers "github.com/LeonardoSola/whatsapp-go-boilerplate/adapters/cmd/whastapp/events"
)

func main() {
	client := whastapp.GetClient()
	fmt.Println("🟢 Connected to WhatsApp")
	defer client.Disconnect()

	client.AddEventHandler(handlers.EventHandler)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
