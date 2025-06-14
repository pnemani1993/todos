package dbutils

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

const SELECT_ALL_STATEMENT string = `SELECT * FROM tasks LIMIT 10;`

func SelectAll(db *sql.DB) ([]DbRow, error) {
	rows, err := db.Query(SELECT_ALL_STATEMENT)
	if err != nil {
		fmt.Println("Query cannot be completed: ", err)
		return nil, err
	}
	var id uint
	var task, description string
	var done, high_priority bool
	var creation_timestamp, alert time.Time
	var dbRow []DbRow = make([]DbRow, 0)
	for rows.Next() {
		err := rows.Scan(&id, &task, &description, &done, &high_priority, &creation_timestamp, &alert)
		if err != nil {
			fmt.Println("Row cannot be obtained: ", err)
			continue
		}
		row := NewResultRow(id, task, description, done,
			high_priority, creation_timestamp, alert)
		dbRow = append(dbRow, row)
	}
	if len(dbRow) == 0 {
		return nil, errors.New("No rows obtained")
	}
	return dbRow, nil
}

func SelectRows(db *sql.DB, selectStatement string) ([]DbRow, error) {
	rows, err := db.Query(selectStatement)
	if err != nil {
		fmt.Println("Query cannot be completed: ", err)
		return nil, err
	}
	var id uint
	var task, description string
	var done, high_priority bool
	var creation_timestamp, alert time.Time
	var dbRow []DbRow = make([]DbRow, 0)
	for rows.Next() {
		err := rows.Scan(&id, &task, &description, &done, &high_priority, &creation_timestamp, &alert)
		if err != nil {
			fmt.Println("Row cannot be obtained: ", err)
			continue
		}
		row := NewResultRow(id, task, description, done,
			high_priority, creation_timestamp, alert)
		dbRow = append(dbRow, row)
	}
	if len(dbRow) == 0 {
		return nil, errors.New("No rows obtained")
	}
	return dbRow, nil
}
