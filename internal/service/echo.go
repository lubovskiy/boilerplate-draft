package service

import (
	"context"

	service "github.com/lubovskiy/test-service/pkg"
	"github.com/sirupsen/logrus"
)

type TestService struct {
	Logger *logrus.Logger
}

func NewTestService() *TestService {
	return &TestService{
		Logger: logrus.New(),
	}
}

func (s *TestService) Echo(ctx context.Context, in *service.EchoRequest) (*service.EchoResponse, error) {
	return &service.EchoResponse{
		Out: in.In,
	}, nil
}
