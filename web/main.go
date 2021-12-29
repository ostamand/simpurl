package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ostamand/url/web/config"
	ctl "github.com/ostamand/url/web/controller"
	"github.com/ostamand/url/web/store/mysql"
)

func main() {
	params := config.Get(os.Getenv("CONFIG_FILE"))
	if _, err := os.Stat("/.dockerenv"); err != nil {
		// running locally not in docker
		params.Db.Addr = "localhost"
	}

	s := mysql.InitializeSQL(&params.Db)
	handler := Handler{storage: s}

	http.HandleFunc("/", handler.redirect)
	http.HandleFunc("/home", handler.home)

	// user
	u := ctl.UserController{Storage: s}
	http.HandleFunc("/signup", u.Signup)
	http.HandleFunc("/signin", u.Signin)
	http.HandleFunc("/signout", u.Signout)

	// links
	l := ctl.LinkController{Storage: s}
	http.HandleFunc("/links/create", l.Create)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("defaulting to port %s", port)
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
