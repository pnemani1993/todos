package dbutils

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// const SELECT_ALL_STATEMENT string = `SELECT * FROM tasks;`

// func SelectAll(db *sql.DB) ([]DbRow, error) {
// 	rows, err := db.Query(SELECT_ALL_STATEMENT)
// 	if err != nil {
// 		fmt.Println("Query cannot be completed: ", err)
// 		return nil, err
// 	}
// 	var id uint
// 	var task, description string
// 	var done, high_priority bool
// 	var creation_timestamp time.Time
// 	var dbRow []DbRow = make([]DbRow, 0)
// 	for rows.Next() {
// 		err := rows.Scan(&id, &task, &description, &done, &high_priority, &creation_timestamp)
// 		if err != nil {
// 			fmt.Println("Row cannot be obtained: ", err)
// 			continue
// 		}
// 		row := NewResultRow(id, task, description, done,
// 			high_priority, creation_timestamp)
// 		dbRow = append(dbRow, row)
// 	}
// 	if len(dbRow) == 0 {
// 		return nil, errors.New("No rows obtained")
// 	}
// 	return dbRow, nil
// }

func selectRows(db *sql.DB, selectStatement string) ([]DbRow, error) {
	rows, err := db.Query(selectStatement)
	if err != nil {
		fmt.Println("Query cannot be completed: ", err)
		return nil, err
	}
	var id uint
	var task, description string
	var done, high_priority bool
	var creation_timestamp time.Time
	var dbRow []DbRow = make([]DbRow, 0)
	for rows.Next() {
		err := rows.Scan(&id, &task, &description, &done, &high_priority, &creation_timestamp)
		if err != nil {
			fmt.Println("Row cannot be obtained: ", err)
			continue
		}
		row := NewResultRow(id, task, description, done,
			high_priority, creation_timestamp)
		dbRow = append(dbRow, row)
	}
	if len(dbRow) == 0 {
		return nil, errors.New("No rows obtained")
	}
	return dbRow, nil
}

func SelectTodoTasks(db *sql.DB) ([]DbRow, error) {
	return selectRows(db, SELECT_NOT_DONE_TASKS)
}

func SelectOnlyDoneTasks(db *sql.DB) ([]DbRow, error) {
	return selectRows(db, SELECT_DONE_TASKS)
}

func SelectHighPriorityTasks(db *sql.DB) ([]DbRow, error) {
	return selectRows(db, SELECT_ALL_HIGH_PRIORITY_TASKS)
}

func SelectAllTasks(db *sql.DB) ([]DbRow, error) {
	return selectRows(db, SELECT_ALL_STATEMENT)
}

func QuerySelector(db *sql.DB, pageName string) ([]DbRow, error) {

	switch pageName {
	case "TodoTasks":
		databaseRow, err := SelectTodoTasks(db)
		if err != nil {
			return nil, err
		}
		return databaseRow, nil

	case "DoneTasks":
		databaseRow, err := SelectOnlyDoneTasks(db)
		if err != nil {
			return nil, err
		}
		return databaseRow, nil

	case "HighPriorityTasks":
		databaseRow, err := SelectHighPriorityTasks(db)
		if err != nil {
			return nil, err
		}
		return databaseRow, nil

	case "AllTasks":
		databaseRow, err := SelectAllTasks(db)
		if err != nil {
			return nil, err
		}
		return databaseRow, nil
	default:
		return nil, errors.New("Invalid value selected")
	}
}

func SelectRow(db *sql.DB) {

}
