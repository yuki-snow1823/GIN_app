package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3" // インポートするけど使わない場合 _ を記述
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

// Model Setting
type Todo struct {
	gorm.Model
	Text   string
	Status string
}

// DB init setting
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("database can't open")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}