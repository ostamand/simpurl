package main

import "github.com/ostamand/url/web/store"

const SessionCookie = "session_token"

func CreateViewModel(user *store.UserModel) *ViewModel {
	loggedIn := false
	if user.ID != 0 {
		loggedIn = true
	}
	return &ViewModel{
		LoggedIn: loggedIn,
		User:     user,
	}
}

type ViewModel struct {
	LoggedIn bool
	User     *store.UserModel
}
