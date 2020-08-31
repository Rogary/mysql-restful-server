package query

import (
	"database/sql"
	"fmt"
	"go-mysql-rest-api/connection"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryDetail(c *gin.Context) {
	if AdminTableFilter(c) {
		queryDetail(c)
	}
}

/**
* func QueryDetail
* table/:id
**/
func queryDetail(c *gin.Context) {
	var (
		result gin.H
	)
	//get param
	table := c.Param("table")
	id := c.Param("id")
	sqlstring := "select * from " + table + " where id = '" + id + "' ;"
	// query

	rows, err := connection.GetConnection().Query(sqlstring)
	defer rows.Close()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, err)
		return
	}
	columns, err := rows.Columns()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, err)
		return
	}
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var m1 map[string]string
	// 再使用make函数创建一个非nil的map. nil不能赋值
	m1 = make(map[string]string)
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusNotFound, err)
			return
		}
		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			m1[columns[i]] = value
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, err)
		return
	}
	result = gin.H{
		"result": m1,
	}
	c.JSON(http.StatusOK, result)
}
