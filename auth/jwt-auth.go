package auth

import (
	"time"

	"go-mysql-rest-api/query"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var jwtMiddleware *jwt.GinJWTMiddleware = nil

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type User struct {
	UserName string
}

var identityKey = "username"

func GetJWTMiddleware() *jwt.GinJWTMiddleware {
	if jwtMiddleware == nil {
		jwtMiddleware = &jwt.GinJWTMiddleware{
			Realm:            "testzone",
			Key:              []byte("secretkey"),
			Timeout:          time.Hour,
			MaxRefresh:       time.Hour,
			SigningAlgorithm: "RS256",
			PrivKeyFile:      "/Users/rogary/Desktop/my/go/src/go-mysql-rest-api/jwtRS256.key",
			PubKeyFile:       "/Users/rogary/Desktop/my/go/src/go-mysql-rest-api/jwtRS256.key.pub",
			IdentityKey:      identityKey,
			PayloadFunc: func(data interface{}) jwt.MapClaims {
				if v, ok := data.(string); ok {
					return jwt.MapClaims{
						identityKey: v,
					}
				}
				return jwt.MapClaims{}
			},
			IdentityHandler: func(c *gin.Context) interface{} {
				claims := jwt.ExtractClaims(c)
				return &User{
					UserName: claims[identityKey].(string),
				}
			},
			Authenticator: func(c *gin.Context) (interface{}, error) {
				var loginVals login
				if err := c.ShouldBind(&loginVals); err != nil {
					return "", jwt.ErrMissingLoginValues
				}

				username := loginVals.Username
				password := loginVals.Password

				if query.CheckUser(username, password) {

					return &User{
						UserName: username,
					}, nil
				}

				return nil, jwt.ErrFailedAuthentication
			},
			Authorizator: func(data interface{}, c *gin.Context) bool {
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
			TokenLookup: "header:Authorization, query: token, cookie: jwt",
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
