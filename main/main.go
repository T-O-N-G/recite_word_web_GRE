package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"regexp"
	"time"
)

const (
	USERNAME = ""
	PASSWORD = ""
	NETWORK  = "tcp"
	SERVER   = ""
	PORT     = 3306
	DATABASE = ""
)

func getJSON(sqlString string, db *sql.DB) (string, error) {
	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(jsonData))
	return string(jsonData), nil
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数

	e.Static("/static", "static")
	e.File("/", "static/index.html")

	e.GET("/word/:word/rand", func(c echo.Context) error {
		word := c.Param("word")
		response, error := getJSON("SELECT * FROM ("+word+") ORDER BY RAND() LIMIT 20", DB)
		if error != nil {
			fmt.Println(err)
		}
		return c.JSON(http.StatusOK, response)
	})

	e.GET("/word/:word/list/:list", func(c echo.Context) error {
		list := c.Param("list")
		word := c.Param("word")
		response, error := getJSON("SELECT * FROM ("+word+") WHERE list=("+list+") ORDER BY RAND()", DB)
		if error != nil {
			fmt.Println(err)
		}
		return c.JSON(http.StatusOK, response)
	})

	e.GET("/word_learn/:word/list/:list", func(c echo.Context) error {
		list := c.Param("list")
		word := c.Param("word")
		response, error := getJSON("SELECT * FROM ("+word+") WHERE WC800.list=("+list+") ORDER BY RAND()", DB)
		if error != nil {
			fmt.Println(err)
		}
		return c.JSON(http.StatusOK, response)
	})

	e.GET("/means_r/:means", func(c echo.Context) error {
		means := c.Param("means")
		pattern := "\\d+" //反斜杠要转义
		result, _ := regexp.MatchString(pattern, means)
		if result == true {
			response, error := getJSON("SELECT mean FROM WC3000 ORDER BY RAND() LIMIT "+means, DB)
			if error != nil {
				fmt.Println(err)
			}
			return c.JSON(http.StatusOK, response)
		} else {
			return c.HTML(http.StatusInternalServerError, "数量必须是个数字")
		}
	})

	e.Logger.Fatal(e.Start(":4000"))

}
