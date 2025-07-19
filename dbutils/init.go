package dbutils

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDatabase() *sql.DB {
	dbConn, err := sql.Open("sqlite3", OsDbLocation())
	if err != nil {
		fmt.Println("Connection cannot be established  ")
		os.Exit(1)
	}
	fmt.Println("Connection established")
	_, err = dbConn.Exec(CREATE_TABLE_LIST)
	if err != nil {
		fmt.Printf("Table cannot be created or accessed %s", err)
		os.Exit(1)
	}
	fmt.Println("tasks table created or already exists")
	return dbConn
}

func OsDbLocation() string {
	os.Setenv("CGO_ENABLED", "1")
	var pathSeparator string
	if runtime.GOOS == "windows" {
		pathSeparator = "\\"
	} else {
		pathSeparator = "/"
	}
	DATABASE_DIR := ".sqlite" + pathSeparator + "databases"
	homeDir, _ := os.UserHomeDir()
	sqlDir := homeDir + pathSeparator + DATABASE_DIR
	fileInfo, err := os.Stat(sqlDir)
	if err != nil || !fileInfo.IsDir() {
		os.Mkdir(sqlDir, 0755)
	}
	sqlDb := sqlDir + pathSeparator + "todos.db"
	return sqlDb
}
