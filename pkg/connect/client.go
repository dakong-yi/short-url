package connect

import (
	"log"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func GetClient() *grpc.ClientConn {
	conn, err := grpc.Dial("127.0.0.1:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return nil
	}
	//defer conn.Close()
	return conn
}
