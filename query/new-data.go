package query

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mysql-rest-api/connection"
	"net/http"
	"strings"
)

func NewData(c *gin.Context) {
	if AdminTableFilter(c) {
		newData(c)
	}
}

func newData(c *gin.Context) {
	table := c.Param("table")
	bodybytes, err := c.GetRawData()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNoContent, gin.H{
			"message": "getBody error",
		})
		return
	}
	m := make(map[string]string)
	err = json.Unmarshal([]byte(bodybytes), &m)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"message": fmt.Sprintf("Unmarshal with error: %+v\n", err),
		})
		return
	}
	keyStrings := ""
	valuesString := ""
	for k, v := range m {
		keyStrings += k
		valuesString += v
		keyStrings += ","
		valuesString += ","
	}

	strings.TrimRight(keyStrings, ",")
	strings.TrimRight(valuesString, ",")

	sqlstring := "INSERT INTO " + table + " ( " + keyStrings + " ) VALUES ( " + valuesString + " );"

	_, err = connection.GetConnection().Exec(sqlstring)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNoContent, gin.H{
			"message": fmt.Sprintf("Exec sql error %s", err.Error()),
		})
	}
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"message": "insert success",
	})

}
