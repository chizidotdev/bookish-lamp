package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/chizidotdev/copia/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(utils.EnvVars.DBDriver, utils.EnvVars.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
