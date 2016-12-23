package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"github.com/labstack/echo"
	//"github.com/labstack/echo/engine/standard"
)
var apiURL string
var err error

func main() {
	app := gin.Default()
	app.LoadHTMLGlob("public/views/*")
	app.StaticFS("/public",http.Dir("./public/"))
	app.StaticFile("/favicon.ico", "public/favicon.ico")
	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
			"static_url" : "http://admin-shorturl.putao.com/",
		})
	})
	app.Run("0.0.0.0:11427")
}
