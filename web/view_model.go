package main

import (
	"net/http"

	"github.com/ostamand/url/web/notify"
	"github.com/ostamand/url/web/store"
)

const SessionCookie = "session_token"

func CreateViewModel(req *http.Request, u *store.UserModel) *ViewModel {
	data := ViewModel{User: u}

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

type ViewModel struct {
	LoggedIn      bool
	User          *store.UserModel
	StatusText    string
	StatusInfo    bool
	StatusSuccess bool
	StatusWarning bool
	StatusError   bool
}
