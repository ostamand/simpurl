package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ostamand/url/web/api"
	"github.com/ostamand/url/web/config"
	ctrl "github.com/ostamand/url/web/controller"
	"github.com/ostamand/url/web/helper"
	"github.com/ostamand/url/web/store/mysql"
)

func main() {
	params := config.Get(os.Getenv("CONFIG_FILE"))

	// hack to run locally thru docker
	if _, err := os.Stat("/.dockerenv"); err != nil {
		params.Db.Addr = "localhost"
	}

	// helpers
	s := mysql.InitializeSQL(&params.Db)
	userHelper := &helper.UserHelper{
		AdminOnly: params.General.AdminOnly,
		Session:   helper.SessionHTTP{},
		Storage:   s,
	}

	// main controller
	h := Handler{
		Storage: s,
		User:    userHelper,
	}
	http.HandleFunc("/", h.redirect)
	http.HandleFunc("/home", h.home)

	// user controller
	u := ctrl.UserController{
		Storage: s,
		User:    userHelper,
	}
	http.HandleFunc("/signup", u.Signup)
	http.HandleFunc("/signin", u.Signin)
	http.HandleFunc("/signout", u.Signout)

	// link controller
	l := ctrl.LinkController{
		Storage: s,
		User:    userHelper,
	}
	http.HandleFunc("/link/create", l.Create)
	http.HandleFunc("/link", l.List)

	// api
	userAPI := api.UserController{Storage: s, User: userHelper}
	http.HandleFunc("/api/signin", userAPI.Signin)

	// file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	// server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
