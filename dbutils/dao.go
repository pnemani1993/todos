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
	Alert             time.Time
}

func NewInsertRow(task string, description string, done bool, highPriority bool, alert time.Time) DbRow {
	var alertTime time.Time = alert.UTC()
	return DbRow{
		Task:         task,
		Description:  description,
		Done:         done,
		HighPriority: highPriority,
		Alert:        alertTime,
	}

}
func NewResultRow(id uint, task string, description string, done bool, highPriority bool, creationTimestamp time.Time, alert time.Time) DbRow {
	creationTimestamp = creationTimestamp.In(time.Local)
	alert = alert.In(time.Local)
	return DbRow{
		Id:                id,
		Task:              task,
		Description:       description,
		Done:              done,
		HighPriority:      highPriority,
		CreationTimestamp: creationTimestamp,
		Alert:             alert,
	}
}
