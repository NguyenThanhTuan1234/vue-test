package main

import (
	"fmt"

	"github.com/labstack/echo"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Test string `json: "test"`
}
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		var task Task
		c.Bind(&task)
		fmt.Println(task.Name)
		fmt.Println(task.Test)
		return nil
	}
}

func main() {
	e := echo.New()

	e.File("/", "index.html")
	e.PUT("/tasks", GetTask())
	e.Logger.Fatal(e.Start(":1234"))
}
