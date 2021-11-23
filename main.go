package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"

	_ "github.com/mattn/go-sqlite3" // インポートするけど使わない場合 _ を記述
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	dbInit()

	// TODO: topと遷移ページ作成したい
	// Bookersっぽくする

	// Todo一覧
	router.GET("/todos", func(ctx *gin.Context) {
		todos := dbGetAllTodo()
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	// Todo追加
	router.POST("/addtodo", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbInsert(text, status)
		// 一覧画面に遷移
		ctx.Redirect(302, "/todos")
	})
	router.Run()

	// Todo詳細画面
	router.GET("/todos/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := dbGetOneTodo(id)
		ctx.HTML(200, "show.html", gin.H{"todo": todo})
	})

	// todo更新
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbUpdate(id, text, status)
		ctx.Redirect(302, "/todos")
	})

	// todo 削除確認
	router.GET("/delete_confirm/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
				panic("todoの削除に失敗しました")
		}
		todo := dbGetOneTodo(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
})
}

// Model 設定
type Todo struct {
	gorm.Model
	Text   string
	Status string
}

// DB init 設定
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
	// Firstの実装確認
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

// delete DB
func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("database can't open :delete")
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

// todo all
func dbGetAllTodo() []Todo {
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("database can't open :getAll")
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

// todo first
func dbGetOneTodo(id int) Todo {
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("database can't open :getOne")
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
