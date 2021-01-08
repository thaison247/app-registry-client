package grpcclient

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"main/protobuf"
)

// validate client app by appId and appToken
func CallValidateClientApp(conn *grpc.ClientConn, clientAppId, clientAppToken, serviceAppId, serviceAppToken, path string) (*protobuf.CommonRes, error) {
	requestMsg := protobuf.ValidateClientReq{
		ClientAppId:    clientAppId,
		ClientAppToken: clientAppToken,
		AppId:          serviceAppId,
		AppToken:       serviceAppToken,
		Path:           path,
	}

	client := protobuf.NewValidateClientServiceClient(conn)

	res, err := client.ValidateClient(context.Background(), &requestMsg)
	if err != nil {
		log.Printf("Call validate client service failed: %v\n", err.Error())
		return nil, err
	}

	return res, err
}
