package main

import (
	"fmt"
	"os"
	"pnemani1993/todos/dbutils"
)

func main() {
	dbConn := dbutils.InitializeDatabase()
	defer dbConn.Close()
	err := dbutils.InsertData(dbConn, dbutils.TRIAL_DATA)
	if err != nil {
		fmt.Println("Data was not inserted")
		os.Exit(1)
	}
	dbRows, err := dbutils.SelectAll(dbConn)
	if err != nil {
		fmt.Println("Data cannot be obtained")
		os.Exit(1)
	}
	for _, row := range dbRows {
		
		fmt.Println("Row: ", row)
	}
	err = dbutils.DeleteData(dbConn, dbutils.DELETE_ALL_TRIAL)
	if err != nil {
		fmt.Println("Data was not deleted")
		os.Exit(1)
	}
	return
}
