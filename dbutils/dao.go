package dbutils

import (
	"time"
)

type DbRow struct {
	Id                uint
	Task              string
	Description       string
	Done              bool
	HighPriority      bool
	CreationTimestamp time.Time
}

func NewInsertRow(task string, description string, done bool, highPriority bool) DbRow {
	return DbRow{
		Task:         task,
		Description:  description,
		Done:         done,
		HighPriority: highPriority,
	}

}
func NewResultRow(id uint, task string, description string, done bool, highPriority bool, creationTimestamp time.Time) DbRow {
	creationTimestamp = creationTimestamp.In(time.Local)
	return DbRow{
		Id:                id,
		Task:              task,
		Description:       description,
		Done:              done,
		HighPriority:      highPriority,
		CreationTimestamp: creationTimestamp,
	}
}

var (
	TRIAL_DATA_ROW []DbRow = []DbRow{
		NewInsertRow("Task 1", "Description for task 1", false, false),
		NewInsertRow("Task 2", "Description for task 2", true, true),
		NewInsertRow("Task 3", "Description for task 3", false, true),
		NewInsertRow("Task 4", "Description for task 4", true, false),
		NewInsertRow("Task 5", "Description for task 5", false, false),
		NewInsertRow("Task 6", "Description for task 6", false, true),
		NewInsertRow("Task 7", "Description for task 7", true, true),
		NewInsertRow("Task 8", "Description for task 8", false, false),
		NewInsertRow("Task 9", "Description for task 9", true, false),
		NewInsertRow("Task 10", "Description for task 10", false, true),
	}
)
