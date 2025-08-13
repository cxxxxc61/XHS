package web

import (
	"github.com/cxxxxc61/XHS/webook/service"
	regexp "github.com/dlclark/regexp2"
)

type UserHandler struct {
	svc          *service.UserService
	csvc         *service.CodeService
	emailtext    *regexp.Regexp
	passwordtext *regexp.Regexp
}

// 预编译正则表达式
func NewUserHandler(svc *service.UserService, csvc *service.CodeService) *UserHandler {
	const (
		emailgrex    = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
		passwordgrex = `^(?=.*[a-zA-Z])(?=.*\d).{1,9}$`
	)
	return &UserHandler{
		svc:          svc,
		csvc:         csvc,
		emailtext:    regexp.MustCompile(emailgrex, regexp.None),
		passwordtext: regexp.MustCompile(passwordgrex, regexp.None),
	}
}
