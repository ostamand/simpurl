package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ostamand/simpurl/cmd/api/controller"
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
		Storage:   s,
	}

	linkAPI := controller.LinkController{Storage: s, User: userHelper}

	http.HandleFunc("/links/create", linkAPI.Create)
	http.HandleFunc("/links", linkAPI.List)
	http.HandleFunc("/links/redirect", linkAPI.Redirect)

	userAPI := controller.UserController{Storage: s, User: userHelper}
	http.HandleFunc("/signin", userAPI.Signin)

	// server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("defaulting to port %s", port)
	}
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
