package service

import "github.com/ax-lew/foaas/domain/model"

type Service interface {
	GetFuckOff(userID string) (*model.Response, error)
}
