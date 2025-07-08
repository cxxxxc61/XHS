package web

import (
	"github.com/cxxxxc61/study/webook/domain"
	"github.com/cxxxxc61/study/webook/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *UserHandler) Signup(c *gin.Context) {
	type Signupreq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirm_password"`
		Password        string `json:"password"`
	}

	var req Signupreq
	//解析从前端拿到的数据
	if err := c.Bind(&req); err != nil {
		return
	}
	//校验邮箱
	ok, err := u.emailtext.MatchString(req.Email)
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		c.String(http.StatusOK, "你的邮箱格式不对")
		return
	}
	//验证密码
	if req.ConfirmPassword != req.Password {
		c.String(http.StatusOK, "两次输入的密码不对")
		return
	}
	//校验密码
	ok, err = u.passwordtext.MatchString(req.Password)
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		c.String(http.StatusOK, "密码需包含字母，数字")
		return
	}

	err = u.svc.Signup(c, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.EmailcomfilctErr {
		c.String(http.StatusOK, "该邮箱已被注册")
		return
	}
	if err != nil {
		c.String(http.StatusOK, "系统异常")
		return
	}
	c.String(http.StatusOK, "注册成功")
	//fmt.Println("&v", req)
}

func (u *UserHandler) Post(c *gin.Context) {
	c.String(http.StatusOK, "post")
}

func (u *UserHandler) Login(c *gin.Context) {
	type Loginreq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Loginreq
	if err := c.Bind(&req); err != nil {
		return
	}
	user, err := u.svc.Login(c, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.PasswordorUserErr {
		c.String(http.StatusOK, "账号/邮箱或密码不对")
		return
	}
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}

	session := sessions.Default(c)
	session.Set("userId", user.Id)
	session.Save()
	c.String(http.StatusOK, "登录成功")

	return
}

func (u *UserHandler) Edit(c *gin.Context) {

}

func (u *UserHandler) Profile(c *gin.Context) {
	c.String(http.StatusOK, "Profile")
}
