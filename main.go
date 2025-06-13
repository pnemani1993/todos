package main

import (
	"pnemani1993/todos/dbutils"
)

func main() {
	dbConn := dbutils.InitializeDatabase()
	defer dbConn.Close()
	return
}
