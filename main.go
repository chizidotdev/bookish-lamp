package main

import (
	"database/sql"
	"github.com/chizidotdev/copia/api"
	db "github.com/chizidotdev/copia/db/sqlc"
	"github.com/chizidotdev/copia/utils"
	"log"
)

func main() {
	utils.LoadConfig()
	log.Println(utils.EnvVars)

	conn, err := sql.Open(utils.EnvVars.DBDriver, utils.EnvVars.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	port := utils.EnvVars.PORT
	if port == "" {
		port = "3333"
	}
	serverAddr := "0.0.0.0:" + port

	err = server.Start(serverAddr)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
