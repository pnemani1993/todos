package dbutils

import (
	"database/sql"
	"fmt"
)

func InsertData(db *sql.DB, insertCommand string) error {
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

func UpdateData(db *sql.DB, updateCommand string) error {
	return nil
}

func DeleteData(db *sql.DB, deleteCommand string) error {
	stmt, err := db.Prepare(deleteCommand)
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

	insertStatement := fmt.Sprintf(`INSERT INTO tasks (task, description, done, high_priority, alert) VALUES
(%s, %s, %d, %d, %s);`, dbRow.Task, dbRow.Description, done, highPriority, dbRow.Alert.Format("2001-01-01 01:00:00"))
	return insertStatement
}
