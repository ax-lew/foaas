package service

import (
	"github.com/ax-lew/foaas/domain/model"
	"github.com/ax-lew/foaas/foaas"
)

type DefaultService struct {
	client *foaas.Client
}

func NewDefaultService(client *foaas.Client) *DefaultService {
	return &DefaultService{client: client}
}

func (s *DefaultService) GetFuckOff(userID string) (*model.Response, error) {
	return s.client.GetFuckOff(userID)
}
