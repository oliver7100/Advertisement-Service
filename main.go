package main

import (
	"net"

	"github.com/oliver7100/advertisement-service/database"
	"github.com/oliver7100/advertisement-service/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")

	if err != nil {
		panic(err)
	}

	d, err := database.NewDatabaseConnection(
		"root:root@tcp(127.0.0.1:3306)/db_user_service?charset=utf8mb4&parseTime=True&loc=Local",
	)

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	proto.RegisterAdvertisementServiceServer(
		s,
		proto.NewService(d),
	)

	panic(s.Serve(listener))
}
