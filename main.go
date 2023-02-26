package main

import (
	"net"

	"github.com/oliver7100/advertisement-service/database"
	"github.com/oliver7100/advertisement-service/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":1338")

	if err != nil {
		panic(err)
	}

	d, err := database.NewDatabaseConnection(
		"services:AVNS_zWWallm_soPdGTwcPQJ@tcp(db-mysql-fmf-do-user-7517862-0.b.db.ondigitalocean.com:25060)/db_advertisement_service",
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
