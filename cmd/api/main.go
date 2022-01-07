package main

import (
	"log"
	"net/http"
	"os"

	ctrl "github.com/ostamand/simpurl/cmd/api/controller"
	"github.com/ostamand/simpurl/internal/config"
	"github.com/ostamand/simpurl/internal/store/mysql"
	"github.com/ostamand/simpurl/internal/user"
)

func main() {
	wd, _ := os.Getwd()
	configPath, _ := config.FindIn(wd, os.Getenv("CONFIG_FILE"))
	params := config.Get(configPath)

	// hack to run locally thru docker
	if _, err := os.Stat("/.dockerenv"); err != nil {
		params.Db.Addr = "localhost"
	}

	// helpers
	s := mysql.InitializeSQL(&params.Db)
	userHelper := &user.UserHelper{
		AdminOnly: params.General.AdminOnly,
		Storage:   s,
	}

	user := ctrl.UserController{Storage: s, User: userHelper}
	http.HandleFunc("/api/signin", user.Signin)

	link := ctrl.LinkController{Storage: s, User: userHelper}
	http.HandleFunc("/api/link/create", link.Create)

	// server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
		log.Printf("defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
