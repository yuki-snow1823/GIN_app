package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	data := "Hello Go/Gin!!"


	router.GET("/", func(ctx *gin.Context) {
		// 第3引数は渡したいデータかも
		// ctx.HTML(200, "index.html", gin.H{})
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	router.Run()
}
