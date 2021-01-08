package grpcclient

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"main/protobuf"
)

// Register APIs for app
func CallApiRegistryService(conn *grpc.ClientConn, appId string, token string, paths []string) (*protobuf.CommonRes, error){
	requestMsg := protobuf.ApiRegistryReq{
		AppId:    appId,
		AppToken: token,
		Path:     paths,
	}

	client := protobuf.NewAppServiceClient(conn)

	res, err := client.ApiRegister(context.Background(), &requestMsg)
	if err != nil {
		log.Printf("Call api registry failed: %v\n", err.Error())
		return nil, err
	}

	return res, err
}