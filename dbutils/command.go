package dbutils

import (
	"database/sql"
	"fmt"
)

func InsertData(db *sql.DB, dbRow DbRow) error {
	insertCommand := fmt.Sprintf(`INSERT INTO tasks (task, description, done, high_priority) VALUES 
					('%s', '%s', %d, %d)`, dbRow.Task, dbRow.Description, isDone(dbRow.Done), isHighPriority(dbRow.HighPriority))
	fmt.Println("Insert command: ", insertCommand)
	stmt, err := db.Prepare(insertCommand)
	if err != nil {
		fmt.Println("Statement cannot be prepared: ", err)
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("Data cannot be inserted: ", err)
		return err
	}
	return nil
}

func isHighPriority(high_priority bool) uint {
	if high_priority {
		return 1
	} else {
		return 0
	}
}

func isDone(done bool) uint {
	if done {
		return 1
	} else {
		return 0
	}
}

func UpdateRow(db *sql.DB, dbRow DbRow) error {
	updateString := fmt.Sprintf(`UPDATE tasks SET description = '%s', done = %d, high_priority = %d  WHERE id = %d`, dbRow.Description, isDone(dbRow.Done), isHighPriority(dbRow.HighPriority), dbRow.Id)
	stmt, err := db.Prepare(updateString)
	if err != nil {
		fmt.Println("Update statement cannot be prepared: ", err)
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("Data cannot be updated: ", err)
		return err
	}
	return nil
}

func SetDone(db *sql.DB, id uint) error {
	updateString := fmt.Sprintf(`UPDATE tasks SET done = 1, high_priority = 0 WHERE id = %d`, id)
	stmt, err := db.Prepare(updateString)
	if err != nil {
		fmt.Println("Task done statement cannot be prepared: ", err)
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("Task done data cannot be updated: ", err)
		return err
	}
	return nil
}

func SetHighPriority(db *sql.DB, id uint) error {
	updateString := fmt.Sprintf(`UPDATE tasks SET high_priority = 1 WHERE id = %d`, id)
	stmt, err := db.Prepare(updateString)
	if err != nil {
		fmt.Println("Task priority setting statement cannot be prepared: ", err)
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("High Priority cannot be updated: ", err)
		return err
	}
	return nil
}

func RemoveHighPriority(db *sql.DB, id uint) error {
	updateString := fmt.Sprintf(`UPDATE tasks SET high_priority = 0 WHERE id = %d`, id)
	stmt, err := db.Prepare(updateString)
	if err != nil {
		fmt.Println("Task priority unsetting statement cannot be prepared: ", err)
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("High Priority cannot be unset: ", err)
		return err
	}
	return nil
}

func DeleteData(db *sql.DB, dbRow DbRow) error {
	delCommand := fmt.Sprintf("Delete from tasks where id = %d", dbRow.Id)
	stmt, err := db.Prepare(delCommand)
	if err != nil {
		fmt.Println("Delete Statement cannot be prepared: ", err)
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("Data cannot be deleted: ", err)
		return err
	}
	return nil
}

func CreateNewRecordStatement(dbRow DbRow) string {
	var done, highPriority uint
	if dbRow.Done {
		done = 1
	} else {
		done = 0
	}
	if dbRow.HighPriority {
		highPriority = 1
	} else {
		highPriority = 0
	}

	insertStatement := fmt.Sprintf(`INSERT INTO tasks (task, description, done, high_priority) VALUES
(%s, %s, %d, %d, %s);`, dbRow.Task, dbRow.Description, done, highPriority)
	return insertStatement
}
