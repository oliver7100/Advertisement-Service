package proto

import (
	"context"

	"github.com/oliver7100/advertisement-service/database"
)

type service struct {
	UnimplementedAdvertisementServiceServer
	Conn *database.Connection
}

func (s *service) CreateAdvertisement(ctx context.Context, req *Advertisement) (*Advertisement, error) {
	m := database.Advertisement{
		Image: req.Image,
		Email: req.Email,
	}

	if tx := s.Conn.Instance.Create(&m); tx.Error != nil {
		return nil, tx.Error
	}

	return req, nil
}

func NewService(conn *database.Connection) *service {
	return &service{
		Conn: conn,
	}
}
