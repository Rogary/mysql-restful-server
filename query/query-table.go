package query

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"go-mysql-rest-api/connection"
)

// QueryList queryList
func QueryList(c *gin.Context) {
	if AdminTableFilter(c) {
		queryList(c)
	}
}

func queryList(c *gin.Context) {
	var (
		results []map[string]string
	)
	table := c.Param("table")
	order := c.DefaultQuery("order", "nil")
	page := c.DefaultQuery("page", "0")
	size := c.DefaultQuery("size", "20")
	pageInt, err := strconv.Atoi(page)
	pageSizeInt, err := strconv.Atoi(size)
	pagePre := strconv.Itoa(pageInt * pageSizeInt)
	pageAfter := strconv.Itoa((pageInt + 1) * pageSizeInt)
	var sqlstring string
	if strings.Compare(sqlstring, "nil") != 0 {
		sqlstring = "select * from " + table + " limit " + pagePre + "," + pageAfter + ";"
	} else {
		sqlstring = "select * from " + table + " order by id " + order + " limit " + pagePre + "," + pageAfter + ";"
	}
	rows, err := connection.GetConnection().Query(sqlstring)
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

	for rows.Next() {
		var m1 map[string]string
		// 再使用make函数创建一个非nil的map. nil不能赋值
		m1 = make(map[string]string)
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

		results = append(results, m1)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusNotFound, err)
			return
		}
	}
	defer rows.Close()
	if results == nil || len(results) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"result": "{}",
			"count":  0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": results,
			"count":  len(results),
		})
	}

}
