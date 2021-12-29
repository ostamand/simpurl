package notify

import (
	"fmt"
	"net/http"
	"strconv"
)

type Notification int
type Status int

const (
	NotifyNone          Notification = 0
	NotifySignedIn      Notification = 101
	NotifyLinkCreated                = 102
	NotifyWrongPassword              = 301
	NotifyNotSignedIn                = 302
)

const (
	StatusNone    Status = 0
	StatusInfo    Status = 1
	StatusSuccess        = 2
	StatusWarning        = 3
	StatusError          = 4
)

var AllNotifications = []Notification{
	NotifyNone, NotifySignedIn,
	NotifyLinkCreated, NotifyWrongPassword, NotifyNotSignedIn,
}

var notificationToText = map[Notification]string{
	NotifyNone:          "",
	NotifySignedIn:      "Welcome! You successfully signed in.",
	NotifyLinkCreated:   "New link created successfully.",
	NotifyWrongPassword: "Your username or password is incorrect. Try again.",
	NotifyNotSignedIn:   "You need to be signed in.",
}

var intToNotification map[int]Notification

func init() {
	intToNotification = make(map[int]Notification)
	for _, n := range AllNotifications {
		intToNotification[int(n)] = n
	}
}

func (n Notification) String() string {
	return notificationToText[n]
}

func (n Notification) Status() Status {
	switch {
	case n > 300:
		return StatusError
	case n > 200:
		return StatusWarning
	case n > 100:
		return StatusSuccess
	case n > 0:
		return StatusInfo
	default:
		return StatusNone
	}
}

func GetNotification(req *http.Request) Notification {
	keys, ok := req.URL.Query()["n"]
	if !ok || len(keys) < 1 {
		return NotifyNone
	}
	if id, err := strconv.Atoi(keys[0]); err == nil {
		return intToNotification[id]
	}
	return NotifyNone
}

func AddNotificationToURL(url string, notification Notification) string {
	return fmt.Sprintf("%s?n=%d", url, int(notification))
}
