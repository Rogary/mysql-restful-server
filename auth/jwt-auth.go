package auth

import (
	"time"

	"go-mysql-rest-api/query"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

var jwtMiddleware *jwt.GinJWTMiddleware = nil

func GetJWTMiddleware() *jwt.GinJWTMiddleware {
	if jwtMiddleware == nil {
		jwtMiddleware = &jwt.GinJWTMiddleware{
			Realm:            "test zone",
			Key:              []byte("secret key"),
			Timeout:          time.Hour,
			MaxRefresh:       time.Hour,
			SigningAlgorithm: "RS256",
			PrivKeyFile:      "testdata/jwtRS256.key",
			PubKeyFile:       "testdata/jwtRS256.key.pub",
			Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
				if query.CheckUser(userId, password) {

					return userId, true
				}

				return userId, false
			},
			Authorizator: func(userId string, c *gin.Context) bool {
				return true
			},
			Unauthorized: func(c *gin.Context, code int, message string) {
				c.JSON(code, gin.H{
					"code":    code,
					"message": message,
				})
			},
			// TokenLookup is a string in the form of "<source>:<name>" that is used
			// to extract token from the request.
			// Optional. Default value "header:Authorization".
			// Possible values:
			// - "header:<name>"
			// - "query:<name>"
			// - "cookie:<name>"
			TokenLookup: "header:Authorization",
			// TokenLookup: "query:token",
			// TokenLookup: "cookie:token",

			// TokenHeadName is a string in the header. Default value is "Bearer"
			TokenHeadName: "Bearer",

			// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
			TimeFunc: time.Now,
		}
	}
	return jwtMiddleware
}
