package notification

import (
	"log"

	"github.com/gen2brain/beeep"
)

func Alert(title, message string) {
	err := beeep.Alert(title, message, "")
	if err != nil {
		log.Printf("Alert fail: %v", err)
	}
}
