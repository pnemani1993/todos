package main

import (
	"pnemani1993/todos/dbutils"
	"pnemani1993/todos/tui"

	"github.com/rivo/tview"
)

func main() {
	trialTui()
	return
}

func trialTui() {
	dbConn := dbutils.InitializeDatabase()
	defer dbConn.Close()
	data := &dbutils.DbRow{}
	tuiApp := tui.TuiApp{Sql: dbConn, App: tview.NewApplication(), Pages: tview.NewPages(), NewTaskPage: tview.NewForm(), Data: data}
	tuiApp.InitApp()
}
