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
	// DBとファイル名
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("database can't open :init")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

// insert DB
func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("database can't open :insert")
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

// update DB
func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("database can't open :update")
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}
