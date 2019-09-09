package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	_ "strings"
)

// Board defines 9x9 sudoku grid
var Board [9][9]int

func init(){
	if true{
		defer fmt.Println(1)
	}
	defer fmt.Println(3)
	defer fmt.Println(2)
}

func main(){
	connectionString := "user=postgres password='babayaga' dbname=gamesapp sslmode=disable" 
	connection, err := sql.Open("postgres", connectionString)
	defer connection.Close()
	if err != nil{
		panic(err)
	}
	if err := connection.Ping(); err!=nil{
		panic(err.Error())
	}
	if stmt, err := connection.Prepare("SELECT * from users"); err == nil{
		defer stmt.Close()
		rows, _ := stmt.Query()
		var (
			a int
			b, c string
		)
		for rows.Next(){
			rows.Scan(&a,&b,&c)
			fmt.Println(a,b,c)
		}
	}
}