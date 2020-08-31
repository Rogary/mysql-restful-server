package query

import (
	"go-mysql-rest-api/conf"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminTableFilter(c *gin.Context) bool {
	table := c.Param("table")
	if strings.Compare(table, conf.GetAuthTableName()) == 0 {
		c.JSON(http.StatusUnauthorized, "{}")
		return false
	}
	return true
}
