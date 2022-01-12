package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ostamand/simpurl/cmd/web/notify"
	"github.com/ostamand/simpurl/internal/store"
	"github.com/ostamand/simpurl/internal/user"
)

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserController struct {
	Storage *store.StorageService
	User    *user.UserHelper
}

func (c *UserController) Signup(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		ShowPage(w, nil, "signup.page.html")
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		u := store.UserModel{
			Username: req.FormValue("username"),
			Password: req.FormValue("password"),
		}
		err := c.Storage.User.Save(&u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
}

func (c *UserController) Signout(w http.ResponseWriter, req *http.Request) {
	c.User.Session.Clear(w)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func (c *UserController) Signin(w http.ResponseWriter, req *http.Request) {
	AllowOrigins(&w)

	switch req.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		data := CreateViewData(req, nil)
		ShowPage(w, data, "signin.page.html")
	case http.MethodPost:
		jsonRequest := req.Header.Get("Content-Type") == "application/json"
		jsonResponse := req.Header.Get("Accept") == "application/json"

		var password, username string
		if jsonRequest{
			request := SigninRequest{}
			body, _ := ioutil.ReadAll(req.Body)
			defer req.Body.Close()

			_ = json.Unmarshal(body, &request)

			username = request.Username
			password = request.Password
		} else {
			if err := req.ParseForm(); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			username = req.FormValue("username")
			password = req.FormValue("password")
		}

		// check password
		u, err := c.User.VerifyPassword(username, password)
		if err != nil || (c.User.AdminOnly && !u.Admin) {
			if jsonResponse {
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else {
				url := notify.AddNotificationToURL("/signin", notify.NotifyWrongPassword)
				http.Redirect(w, req, url, http.StatusSeeOther)
				return
			}
		}

		// create session token
		session, err := c.User.CreateSession(w, u.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if jsonResponse {
			data := struct {
				Token string `json:"token"`
			}{
				Token: session.Token,
			}
			json.NewEncoder(w).Encode(data)
		} else {
			http.Redirect(
				w,
				req,
				notify.AddNotificationToURL("/home", notify.NotifySignedIn),
				http.StatusSeeOther,
			)
		}
	}
}
