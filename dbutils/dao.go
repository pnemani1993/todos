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
