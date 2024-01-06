package notification_test

import (
	"testing"

	"github.com/chrisngyn/torplayer/notification"
)

func TestAlert(t *testing.T) {
	notification.Alert("Test", "This is a test")
}
