package main

import (
	_ "todoList/docs"
	"todoList/internal/app"
)

// @title Swagger TodoList API
// @version 1.0
// @description This is a Simple TodoList.

// @contact.name Daniil Koreshkov
// @contact.url https://t.me/Kr4ckeT
// @contact.email danielkoreshkov@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func main() {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
