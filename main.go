package main

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

type H map[string]interface{}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "test"
	SSL         = "disable"
)

type User struct {
	UserID   int
	Name     string
	Age      int `json:",string"`
	Location string
}
type Users struct {
	Users []User `json:"users"`
}

// var DB *sql.DB
var err error

func DBinit() *sql.DB {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		DB_USER, DB_PASSWORD, DB_NAME, SSL)
	DB, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success")
	return DB
}
func GetUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user User

		c.Bind(&user)
		fmt.Println(user.Name)
		fmt.Println(user.Age)
		fmt.Println(user.Location)
		// age, err := strconv.Atoi(user.Age)
		// if err != nil {
		// 	panic(err)
		// }
		// id, err := PutUser(db, user.Name, age, user.Location)
		// if err == nil {
		// 	return c.JSON(http.StatusCreated, H{
		// 		"created": id,
		// 	})
		// 	// Handle any errors
		// } else {
		// 	return err
		// }
		return nil
	}
}

func PutUser(db *sql.DB, name string, age int, location string) (int64, error) {
	sql := "INSERT INTO users(name, age, location) VALUES($1, $2, $3)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(name, age, location)
	if err != nil {
		panic(err)
	}
	return result.LastInsertId()
}

func main() {
	e := echo.New()
	DB := DBinit()
	e.File("/", "public/index.html")
	e.Static("/*", "public")
	e.Static("/css/*", "public/css")
	e.Static("/fonts/*", "public/fonts")
	e.Static("/js/*", "public/js")
	e.Static("/images/*", "public/images")
	e.Static("/sass/*", "public/sass")
	e.File("/users", "index.html")
	e.POST("/users", GetUser(DB))
	e.Logger.Fatal(e.Start(":1234"))
}
