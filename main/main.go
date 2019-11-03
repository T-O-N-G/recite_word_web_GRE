package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
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

func getJSON(sqlString string, db *sql.DB, query string) (string, error) {
	rows, err := db.Query(sqlString, query)
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
	fmt.Println(string(jsonData))
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

	e.GET("/", func(c echo.Context) error {
		if err != nil {
			fmt.Println(err)
		}
		response, error := getJSON("SELECT * FROM WC1500 ORDER BY RAND() LIMIT ?", DB, "20")
		if error != nil {
			fmt.Println(err)
		}
		return c.JSON(http.StatusOK, response)
	})

	e.GET("/list/:list", func(c echo.Context) error {
		list := c.Param("list")
		fmt.Println(list)
		response, error := getJSON("SELECT * FROM WC1500 WHERE WC1500.list=?", DB, list)
		if error != nil {
			fmt.Println(err)
		}
		return c.JSON(http.StatusOK, response)
	})

	e.GET("/means_r/:means", func(c echo.Context) error {
		means := c.Param("means")
		fmt.Println(means)
		response, error := getJSON("SELECT mean FROM WC1500 ORDER BY RAND() LIMIT ?", DB, means)
		if error != nil {
			fmt.Println(err)
		}
		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":4000"))
}
