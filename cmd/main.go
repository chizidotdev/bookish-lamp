package main

import (
	"database/sql"
	"github.com/chizidotdev/copia/internal/app"
	"github.com/chizidotdev/copia/internal/repository/sqlx"
	"github.com/chizidotdev/copia/internal/service"
	"github.com/chizidotdev/copia/pkg/utils"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	utils.LoadConfig()

	conn, err := sql.Open(utils.EnvVars.DBDriver, utils.EnvVars.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := sqlx.NewStore(conn)
	newService := service.NewService(store)
	server := app.NewServer(newService)

	port := utils.EnvVars.PORT
	if port == "" {
		port = "8080"
	}
	serverAddr := "0.0.0.0:" + port

	err = server.Start(serverAddr)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
