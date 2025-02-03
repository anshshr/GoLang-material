package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("learning about go to make the servers in go and do get and the post request")

	route := gin.Default() // setting the gin route

	route.GET("/" , func(ctx *gin.Context) { // ctx - stores the data of req amd res from the server
		ctx.JSON(200 , gin.H{
			"message" : "now this is the start of the golang",
		})
	})


	route.GET("/launch/:url" , func(ctx *gin.Context) { // before the dynamic routing we had to placea prior an route before here ?:ulr will not work but /data/:url is working
		name := ctx.Param("url")
		fmt.Println(name);

		// ctx.JSON(200 , gin.H{
		// 	"url" : name,
		// })

		ctx.Redirect(http.StatusFound , "https://www.youtube.com/watch?v=Vx2zPMPvmug&t=940s")


	})

	route.Run(":8080")

}
