package web

import (
	"github.com/cxxxxc61/study/webook/internal/web/middleware"
	"github.com/cxxxxc61/study/webook/repository"
	"github.com/cxxxxc61/study/webook/repository/dao"
	"github.com/cxxxxc61/study/webook/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

type UserHandler struct {
	svc          *service.UserService
	emailtext    *regexp.Regexp
	passwordtext *regexp.Regexp
}

// 预编译正则表达式
func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailgrex    = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
		passwordgrex = `^(?=.*[a-zA-Z])(?=.*\d).{1,9}$`
	)
	return &UserHandler{
		svc:          svc,
		emailtext:    regexp.MustCompile(emailgrex, regexp.None),
		passwordtext: regexp.MustCompile(passwordgrex, regexp.None),
	}
}

func initdb() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:BazKT%HlbsP3@tcp(117.50.198.118:30336)/cxc_webook"))
	if err != nil {
		panic(err)
	}
	err = dao.Inittable(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initUser(db *gorm.DB) *UserHandler {
	d := dao.NewUserDao(db)
	repo := repository.NewUserRepository(d)
	svc := service.NewUserService(repo)
	u := NewUserHandler(svc)
	return u
}

func initwebsever() *gin.Engine {
	sever := gin.Default()
	sever.Use(cors.New(cors.Config{
		//详细地址
		//AllowOrigins:     []string{"https://foo.com"},
		AllowMethods: []string{"POST", "GET"},
		//业务请求中可以带上的头
		AllowHeaders: []string{"Origin"},
		//允许传入的头
		ExposeHeaders: []string{"Content-Length"},
		//用户认证信息
		AllowCredentials: true,
		//类型地址
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "your.com")
		},
		MaxAge: 12 * time.Minute,
	}))

	store, err := redis.NewStore(64,
		"tcp", "117.50.198.118:30379",
		"", "&NCNOiJdzC79",
		[]byte("bHO2mkqCDKSB2GsqikJGlQURD0KtwiuZI4zpWZYolG7QCE64hTM0r6O5VhrdjFHt"))
	if err != nil {
		panic(err)
	}
	sever.Use(sessions.Sessions("mysession", store))
	sever.Use(middleware.NewLoginMiddlewareBuild().
		Ignorepath("/users/login").
		Ignorepath("/users/signup").Build())

	return sever
}
