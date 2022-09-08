package service

import "context"

type Service interface {
	GetUserInfo(ctx context.Context, token string) (interface{}, error)
}

func New() Service {
	var srv Service
	srv = NewBasicService()
	return srv
}

func NewBasicService() Service {
	return basicService{}
}

type basicService struct{}

func (b basicService) GetUserInfo(ctx context.Context, token string) (interface{}, error) {

}
