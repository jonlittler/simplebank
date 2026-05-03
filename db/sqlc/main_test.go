package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:password@localhost:5432/simplebank?sslmode=disable"
)

var testStore Store

func TestMain(m *testing.M) {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	testStore = NewStore(conn)
	os.Exit(m.Run())
}
