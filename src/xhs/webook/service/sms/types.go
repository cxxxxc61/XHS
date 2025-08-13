package sms

import (
	"context"
)

type Service interface {
	Send(c context.Context, tql string,
		args []string, numbers ...string) error
}
