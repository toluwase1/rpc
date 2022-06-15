package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"rpc-client-server/models"
)

func main() {
	connection, err := net.Dial("tcp", ":3000")
	if err != nil {
		log.Fatalf("Could not connect to localhost:3000. %s", err)
	}
	defer connection.Close()
	client := jsonrpc.NewClient(connection)
	/* Test Mining.Authorize */
	var authorizationResponse bool
	if err := client.Call("Mining.Authorize", models.AuthorizeRequest{Username: "toluwase", Password: "toluwase"}, &authorizationResponse); err != nil {
		log.Fatalf("An error occurred, %s", err)
	}
	fmt.Printf("Mining.Authorize response: %v", authorizationResponse)
	/*Test Mining.Subscribe */
	var subscribeResponse models.SubscribeResponse
	if err := client.Call("Mining.Subscribe", "", &subscribeResponse); err != nil {
		log.Fatalf("An error occurred: %s", err)
	}
	fmt.Println("Mining.Subscribe response: ", subscribeResponse)

}
