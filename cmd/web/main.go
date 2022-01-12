package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	ctrl "github.com/ostamand/simpurl/cmd/web/controller"
	"github.com/ostamand/simpurl/cmd/web/controller/api"
	"github.com/ostamand/simpurl/cmd/web/session"
	"github.com/ostamand/simpurl/internal/config"
	"github.com/ostamand/simpurl/internal/store/mysql"
	"github.com/ostamand/simpurl/internal/user"
)

func main() {
	wd, _ := os.Getwd()
	configPath, _ := config.FindIn(wd, os.Getenv("CONFIG_FILE"))
	fmt.Println(configPath)
	params := config.Get(configPath)

	// hack to run locally thru docker
	if _, err := os.Stat("/.dockerenv"); err != nil {
		params.Db.Addr = "localhost"
	}

	// helpers
	s := mysql.InitializeSQL(&params.Db)
	userHelper := &user.UserHelper{
		AdminOnly: params.General.AdminOnly,
		Session:   &session.SessionHTTP{},
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
	linkAPI := api.LinkController{Storage: s, User: userHelper}
	http.HandleFunc("/api/links/create", linkAPI.Create)

	// file server
	fileServer := http.FileServer(http.Dir("./web/static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	// server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
