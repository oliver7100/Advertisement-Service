package proto

import (
	"context"
	"encoding/json"

	"github.com/oliver7100/advertisement-service/database"
)

type service struct {
	UnimplementedAdvertisementServiceServer
	Conn *database.Connection
}

func TypeConverter[R any](data any) (*R, error) {
	var result R
	b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (s *service) CreateAdvertisement(ctx context.Context, req *Advertisement) (*Advertisement, error) {
	m, err := TypeConverter[database.Advertisement](req)

	if err != nil {
		return nil, err
	}

	if tx := s.Conn.Instance.Create(&m); tx.Error != nil {
		return nil, tx.Error
	}

	return req, nil
}

func (s *service) DisableAdvertisement(ctx context.Context, req *DisableAdvertisementRequest) (*Advertisement, error) {
	var o database.Advertisement
	if tx := s.Conn.Instance.Model(&database.Advertisement{ID: int(req.Id)}).First(&o); tx.Error != nil {
		return nil, tx.Error
	}

	o.Activated = false

	if tx := s.Conn.Instance.Save(o); tx.Error != nil {
		return nil, tx.Error
	}

	return &Advertisement{
		Activated: o.Activated,
	}, nil
}

func (s *service) GetAdvertisements(ctx context.Context, req *GetAllAdvertisementsRequest) (*GetAllAdvertisementsResponse, error) {
	var items []*Advertisement

	if tx := s.Conn.Instance.Table("advertisements").Find(&items, "activated = ?", 0); tx.Error != nil {
		return nil, tx.Error
	}

	return &GetAllAdvertisementsResponse{
		Items: items,
	}, nil
}

func NewService(conn *database.Connection) *service {
	return &service{
		Conn: conn,
	}
}
