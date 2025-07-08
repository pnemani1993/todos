package main

import (
	"fmt"
	"os"
	"pnemani1993/todos/dbutils"
	"pnemani1993/todos/tui"
)

func main() {
	// tui.InitPage()
	// tui.InitButtons()
	tui.GetInitModal()
	// tui.TryingPages()
	// trialRun()
	return
}

func trialRun() {
	dbConn := dbutils.InitializeDatabase()
	defer dbConn.Close()
	for _, row := range dbutils.TRIAL_DATA_ROW {
		err := dbutils.InsertData(dbConn, row)
		if err != nil {
			fmt.Println("Data was not inserted: ", row)
			// os.Exit(1)
		}
	}
	fmt.Println("Data insertion successful")
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
	fmt.Println("The trial ended.")
}
