package whastapp

import (
	"go.mau.fi/whatsmeow"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var client *whatsmeow.Client

func init() {
	clientLog := waLog.Stdout("Client", "ERROR", true)
	deviceStore := getStorage()

	client = whatsmeow.NewClient(deviceStore, clientLog)

	var err error
	if client.Store.ID == nil {
		err = Connect()
	} else {
		err = client.Connect()
	}

	if err != nil {
		panic(err)
	}
}

func GetClient() *whatsmeow.Client {
	return client
}
