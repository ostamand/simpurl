package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ostamand/url/web/config"
	ctrl "github.com/ostamand/url/web/controller"
	"github.com/ostamand/url/web/store/mysql"
	"github.com/ostamand/url/web/user"
)

func main() {
	params := config.Get(os.Getenv("CONFIG_FILE"))

	// hack to run locally thru docker
	if _, err := os.Stat("/.dockerenv"); err != nil {
		params.Db.Addr = "localhost"
	}

	// helpers
	s := mysql.InitializeSQL(&params.Db)
	userHelper := user.UserHelper{
		AdminOnly:      params.General.AdminOnly,
		StorageService: s,
	}

	// main controller
	h := Handler{
		StorageService: s,
		User:           &userHelper,
	}
	http.HandleFunc("/", h.redirect)
	http.HandleFunc("/home", h.home)

	// user controller
	u := ctrl.UserController{
		StorageService: s,
		User:           &userHelper,
	}
	http.HandleFunc("/signup", u.Signup)
	http.HandleFunc("/signin", u.Signin)
	http.HandleFunc("/signout", u.Signout)

	// link controller
	l := ctrl.LinkController{
		StorageService: s,
		User:           &userHelper,
	}
	http.HandleFunc("/link/create", l.Create)
	http.HandleFunc("/link", l.List)

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
