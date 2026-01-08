package database

import "fmt"

var connection string

func init() {
	fmt.Println("Ini init")
	connection = "MySQL"
}

func GetData() string {
	return connection
}