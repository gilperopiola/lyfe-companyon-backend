package main

import (
	"flag"
	"log"
	"os"

	"github.com/gilperopiola/lyfe-companyon-backend/config"
	"github.com/gilperopiola/lyfe-companyon-backend/database"
)

var cfg config.MyConfig
var db database.MyDatabase
var rtr MyRouter

func main() {
	env := flag.String("env", "", "local | dev | prod")
	flag.Parse()

	if *env == "" {
		log.Fatal("environment flag not set")
	}

	cfg.Setup(*env)
	db.Setup(cfg)
	defer db.Close()
	rtr.Setup()

	log.Println("server started")
	rtr.Run(":" + os.Getenv("PORT"))
}
