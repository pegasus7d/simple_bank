package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/pegasus7d/simplebank/db/util"
)

var testQueries *Queries

var testDB *sql.DB
var err error


func TestMain(m *testing.M){
	config,err:=util.LoadConfig("../..")
	if err!=nil{
		log.Fatal("Cannot load config files",err)
	}
	testDB,err=sql.Open(config.DBDriver,config.DBSource)
	if err !=nil{
		log.Fatal("Cannot connect to db:",err)
	}
	testQueries=New(testDB)
	os.Exit(m.Run())
}
