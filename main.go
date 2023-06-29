package main

import (
	"database/sql"
	"log"

	"github.com/chizidotdev/copia/api"
	db "github.com/chizidotdev/copia/db/sqlc"
	"github.com/chizidotdev/copia/utils"
)

func init() {

	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatal("Cannot load config:", err)
	//}
	utils.LoadConfig()
}

func main() {
	conn, err := sql.Open(utils.EnvVars.DBDriver, utils.EnvVars.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(utils.EnvVars.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
