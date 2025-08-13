package memory

import (
	"context"
	"fmt"
)

type Service struct {
}

func (s *Service) Send(c context.Context, tql string, args []string, numbers ...string) error {
	fmt.Println(args)
	return nil
}

func NewsmsService() *Service {
	return &Service{}
}
