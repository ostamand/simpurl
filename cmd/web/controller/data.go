package controller

import (
	"net/http"

	"github.com/ostamand/url/cmd/web/notify"
	"github.com/ostamand/url/internal/store"
)

const SessionCookie = "session_token"

func CreateViewData(req *http.Request, u *store.UserModel) *ViewData {
	data := ViewData{User: u}

	// flag for user login
	if u != nil && u.ID != 0 {
		data.LoggedIn = true
	}

	// process notifications from url
	n := notify.GetNotification(req)
	if n != notify.NotifyNone {
		switch n.Status() {
		case notify.StatusInfo:
			data.StatusInfo = true
		case notify.StatusSuccess:
			data.StatusSuccess = true
		case notify.StatusWarning:
			data.StatusWarning = true
		case notify.StatusError:
			data.StatusError = true
		}
		data.StatusText = n.String()
	}

	return &data
}

type ViewData struct {
	LoggedIn      bool
	User          *store.UserModel
	StatusText    string
	StatusInfo    bool
	StatusSuccess bool
	StatusWarning bool
	StatusError   bool
}
