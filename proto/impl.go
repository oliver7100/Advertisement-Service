package proto

import "context"

type IService interface {
	CreateAdvertisement(context.Context, *Advertisement) (*Advertisement, error)
}
