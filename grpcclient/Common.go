package grpcclient

import (
	"google.golang.org/grpc"
	"log"
)

// get grpc connection to server
func GetConnection(target string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Printf("Connection failed: %v\n", err.Error())
		return nil, err
	}

	return conn, err
}