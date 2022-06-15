package main

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"rpc-client-server/database"
	"rpc-client-server/handlers"
)

func main() {
	/* Load configuration */
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Could not load app configuration. %s", err)
	}

	/* Database connection */
	database.InitPostgresDB(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	defer func(Db *gorm.DB) {
		err := Db.Close()
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}(database.Db)

	rpcServer := rpc.NewServer()
	mining := new(handlers.Mining)
	err := rpcServer.Register(mining)
	if err != nil {
		log.Fatalf("Could not register service, %s", err)
	}

	rpcServer.HandleHTTP("/", "/debug")
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal("Could not listen on port 3000")
	}
	log.Println("Started RPC Handler on localhost:3000")

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		codec := jsonrpc.NewServerCodec(connection)
		go rpcServer.ServeCodec(codec)
	}

}
