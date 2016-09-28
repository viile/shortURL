package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	app := echo.New()
	app.GET("/user/:id", func(c echo.Context) error {
		id := c.Param("id")
		fmt.Println("params : ",id)
		response,_ := http.Get("http://children-account.putao.com/u?keyword=" + id)
		defer response.Body.Close()
		body,_ := ioutil.ReadAll(response.Body)
		fmt.Println("result : ",string(body))
		return c.String(http.StatusOK, string(body))
	})
	fmt.Println("open http://127.0.0.1:11427/user/keyword in your browser")
	fmt.Println("eg: http://127.0.0.1:11427/user/60060427")
	fmt.Println("eg: http://127.0.0.1:11427/user/18516560427")
	fmt.Println("eg: http://127.0.0.1:11427/user/mmaabbcc0427@qq.com")
	app.Run(standard.New("0.0.0.0:11427"))
}
