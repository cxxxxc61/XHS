package service

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/cxxxxc61/XHS/webook/repository"
	"github.com/cxxxxc61/XHS/webook/service/sms"
)

const codetplid = "1111111"

type CodeService struct {
	repo   *repository.CodeRepository
	smssvc sms.Service
}

func NewCodeService(repo *repository.CodeRepository, smssvc sms.Service) *CodeService {
	return &CodeService{
		repo:   repo,
		smssvc: smssvc,
	}
}

func (csvc *CodeService) Send(c context.Context,
	biz string,
	phone string) error {
	code := csvc.genrateCode()
	err := csvc.repo.Store(c, code, biz, phone)
	if err != nil {
		return err
	}
	err = csvc.smssvc.Send(c, codetplid, []string{code}, phone)
	return err
}

func (csvc *CodeService) Verify(c context.Context,
	biz string,
	phone string,
	inputcode string) (bool, error) {
	return csvc.repo.Verify(c, biz, phone, inputcode)
}

func (csvc *CodeService) genrateCode() string {
	num := rand.Intn(1000000)
	return fmt.Sprintf("%06d", num)
}
