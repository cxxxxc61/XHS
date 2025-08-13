package middleware

import (
	"github.com/cxxxxc61/XHS/webook/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

type LoginjwtMiddlewareBuild struct {
	path []string
}

func NewLoginjwtMiddlewareBuild() *LoginjwtMiddlewareBuild {
	return &LoginjwtMiddlewareBuild{}
}

func (LMB *LoginjwtMiddlewareBuild) Ignorepath(path string) *LoginjwtMiddlewareBuild {
	LMB.path = append(LMB.path, path)
	return LMB
}

func (LMB *LoginjwtMiddlewareBuild) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if c.Request.URL.Path == "/users/login" ||
		//	c.Request.URL.Path == "/users/signup" {
		//	return
		//}
		//
		for _, path := range LMB.path {
			if path == c.Request.URL.Path {
				return
			}
		}
		tokenheader := c.GetHeader("Authorization")
		if tokenheader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segs := strings.Split(tokenheader, " ")
		if len(segs) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		tokenstr := segs[1]
		claims := &web.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenstr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("bHO2mkqCDKSB2GsqikJGlQURD0KtwiuZI4zpWZYolG7QCE64hTM0r6O5VhrdjFHt"), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims.ExpiresAt.Sub(time.Now()) < time.Second*50 {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
			tokenstr, err = token.SignedString([]byte("bHO2mkqCDKSB2GsqikJGlQURD0KtwiuZI4zpWZYolG7QCE64hTM0r6O5VhrdjFHt"))
			if err != nil {
				log.Println("jwt续约失败", err)
			}
			c.Header("x-jwt-token", tokenstr)
		}

		c.Set("claims", claims)
	}

}
