package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"rpc-client-server/models"
)

func main() {
	log.Println("hello 1")
	connection, err := net.Dial("tcp", ":3000")
	log.Println("hello 2")
	if err != nil {
		log.Fatalf("Could not connect to localhost:3000. %s", err)
	}
	log.Println("hello 3")
	defer connection.Close()
	log.Println("hello 4")
	client := jsonrpc.NewClient(connection)
	log.Println("hello 5")
	/* Test Mining.Authorize */
	log.Println("hello 6")
	var authorizationResponse bool
	if err := client.Call("Mining.Authorize", models.AuthorizeRequest{Username: "toluwase", Password: "toluwase"}, &authorizationResponse); err != nil {
		log.Fatalf("An error occurred, %s", err)
	}
	fmt.Printf("Mining.Authorize response: %v", authorizationResponse)
	fmt.Println()
	log.Println("hello 7")
	/*Test Mining.Subscribe */
	var subscribeResponse models.SubscribeResponse
	log.Println("hello 8")
	if err := client.Call("Mining.Subscribe", "", &subscribeResponse); err != nil {
		log.Fatalf("An error occurred: %s", err)
	}
	fmt.Println("Mining.Subscribe response: ", subscribeResponse)

}
