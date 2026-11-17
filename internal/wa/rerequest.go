package wa

import (
	"os"
	"strings"
)

// RerequestFromPhone says whether a message this device could not decrypt is
// asked back from the primary phone. It is on unless WACLI_REREQUEST_FROM_PHONE
// is 0, false, no or off. Each answer from the phone can show the person a
// "sync completed" notification (seen on iPhone), so a program that keeps
// sync running may prefer the gap to the notifications.
func RerequestFromPhone() bool {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("WACLI_REREQUEST_FROM_PHONE"))) {
	case "0", "false", "no", "off":
		return false
	}
	return true
}
