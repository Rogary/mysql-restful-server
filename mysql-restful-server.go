package main

import (
	"go-mysql-rest-api/auth"
	"go-mysql-rest-api/conf"
	"go-mysql-rest-api/connection"
	"go-mysql-rest-api/query"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//create mysql connection
	initPool()
	router := gin.Default()
	// GET a detail
	authMiddleware := auth.GetJWTMiddleware()
	router.GET("/api/v1/:table/:id", query.QueryDetail)
	router.GET("/api/v1/:table", query.QueryList)
	router.POST("/login", authMiddleware.LoginHandler)
	auths := router.Group("/api")
	auths.Use(authMiddleware.MiddlewareFunc())
	{
		auths.DELETE("/v1/:table/:id", query.DeleteDetail)
		auths.POST("/v1/:table", query.NewData)
		auths.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	router.Run(":8989")

}
func initPool() {
	connection.InitConnection(conf.GetMysqlDataSourc())
}
