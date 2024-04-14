package main

import (
	"log"

	"github.com/mfarrasml/template-authorization-app/config"
	"github.com/mfarrasml/template-authorization-app/database"
	"github.com/mfarrasml/template-authorization-app/router"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.InitDb(config.DbUrl())
	if err != nil {
		log.Fatal("error connecting to db")
	}
	defer db.Close()

	err = router.ServeRouter(db, *config)
	if err != nil {
		log.Fatal(err)
	}
}
