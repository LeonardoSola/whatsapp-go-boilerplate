package whastapp

import (
	"context"
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"
)

func Connect() error {
	qrChan, _ := client.GetQRChannel(context.Background())
	err := client.Connect()
	if err != nil {
		panic(err)
	}

	tries := 3

	for evt := range qrChan {
		switch evt.Event {
		case "code":
			qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
		case "success":
			return nil
		case "timeout":
			tries--
			if tries > 0 {
				break
			}
		}
	}

	return fmt.Errorf("login failed")
}
