package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"main/protobuf"
)



func main() {
	fmt.Println("test main func")

	log.Println("Client running ...")

	conn, err := grpc.Dial(":9090", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	validateClient := protobuf.NewValidateClientServiceClient(conn)

	validateReq := protobuf.ValidateClientReq{
		ClientAppId:    "APP004",
		ClientAppToken: "TYXARUH82QQ9T7FQWBL8",
		AppId:          "APP002",
		AppToken:       "BEKP9APS3NJEYRJNXKBA",
		Path:           "/new/url",
	}

	res, err := validateClient.ValidateClient(context.Background(), &validateReq)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Println(res)

}


func GetClientConnection(target string) (*grpc.ClientConn, error){
	conn, err := grpc.Dial(":9090", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	return conn, err
}
