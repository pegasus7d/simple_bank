package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/pegasus7d/simplebank/api"
	db "github.com/pegasus7d/simplebank/db/sqlc"
	"github.com/pegasus7d/simplebank/db/util"
)

const (
	dbDriver="postgres"
	dbSource="postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress="0.0.0.0:8080"
)



func main(){
	config,err:=util.LoadConfig(".")
	if err!=nil{
		log.Fatal("Cannot load config files",err)
	}
	conn,err:=sql.Open(config.DBDriver,config.DBSource)
	if err !=nil{
		log.Fatal("Cannot connect to db:",err)
	}


	store:=db.NewStore(conn)
	server:=api.NewServer(store)

	err=server.Start(config.ServerAddress)
	if err!=nil{
		log.Fatal("cannot start server",err)
	}

}