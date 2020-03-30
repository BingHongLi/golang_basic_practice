package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Demo struct{

	DemoString string `json:"demoString" form:"demoString" binding:"required"`
	DemoInt int `json:"demoInt" form:"demoString" `

}

func main() {


	// 準備一個gin server，並調用預設的中間層
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// 設定一個最經典的功能，
	// 開發者可用get訪問 localhost:8080/hello
	// 取得基礎回應
	// basic
	router.GET("/hello",func(c *gin.Context){
		c.String(200, "Hello World")
	})

	// 在HTTP的使用場景裡，有時會透過URL Pararmeter當作對系統的輸入
	// 開發者可用 get訪問  localhost:8080/user/{variable}
	// 系統可接收此Variable，並處理
	// just for url pararmeter
	// https://github.com/gin-gonic/gin#bind-uri
	router.GET("/user/:name", func(c *gin.Context){

		userName := c.Param("name")
		c.String(200,userName)

	})

	// 在HTTP的系統裡，經常使用GET 與 POST 這兩個動作 對系統進行訪問
	// 使用GET的時候，經常會透過Query String，給系統下達指令
	// 使用POST的時候，經常會透過 封包裡的內容 給系統下達指令
	// 此處展示如何parse出 QueryString 與 Request Body
	// BindPostJson or QueryString
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	router.GET("/demo",func(c *gin.Context){
		var demoFromQueryString Demo
		c.ShouldBind(&demoFromQueryString)
		c.JSON(200,demoFromQueryString)
	})
	router.POST("/demo",func(c *gin.Context){
		var demoFromQueryString Demo
		c.ShouldBind(&demoFromQueryString)
		c.JSON(200,demoFromQueryString)
	})


	// 進階
	// 使用網路服務的時候，有時候會透過Cookie，做系統上的安全憑證
	// gin操作Cookie的方式如下
	// https://github.com/gin-gonic/gin#set-and-get-a-cookie
	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("cxcxc_auth_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("cxcxc_auth_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})


	// 若要額外提供圖片等檔案做輸出
	// https://github.com/gin-gonic/gin#serving-static-files
	router.Static("/static", "./static")

	// 應用場景
	// https://github.com/gin-gonic/gin#serving-data-from-file

	// 應用場景
	// https://github.com/gin-gonic/gin#try-to-bind-body-into-different-structs

	// 應用場景
	// https://github.com/gin-gonic/gin#http2-server-push

	// 應用場景
	// https://github.com/gin-gonic/gin#testing

	router.Run()

}
