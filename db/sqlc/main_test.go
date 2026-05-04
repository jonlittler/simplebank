package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/jonlittler/ts/simplebank/util"
	_ "github.com/lib/pq"
)

var testStore Store

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	testStore = NewStore(conn)
	os.Exit(m.Run())
}
