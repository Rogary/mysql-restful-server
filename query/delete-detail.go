package query

import (
	"fmt"
	"go-mysql-rest-api/connection"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteDetail(c *gin.Context) {
	if AdminTableFilter(c) {
		deleteDetail(c)
	}
}

/**
* func DeleteDetail
* table/:id
**/
func deleteDetail(c *gin.Context) {
	//get param
	table := c.Param("table")
	id := c.Param("id")
	sqlstring := "delete  from " + table + " where id = '" + id + "' ;"
	// query

	result, err := connection.GetConnection().Exec(sqlstring)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNoContent, gin.H{
			"message": fmt.Sprintf("%s not found", id),
		})
	} else {
		ids, erro := result.RowsAffected()
		if erro != nil || ids == 0 {
			fmt.Println(erro)
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Failed deleted user: %s ", id),
			})
		} else {

			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Successfully deleted user: %s", id),
			})
		}
	}

}
