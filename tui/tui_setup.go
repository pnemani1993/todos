package tui

import (
	"database/sql"
	"fmt"
	"pnemani1993/todos/dbutils"

	"time"

	"github.com/rivo/tview"
)

type TuiApp struct {
	Sql         *sql.DB
	App         *tview.Application
	Pages       *tview.Pages
	NewTaskPage *tview.Form
	Data        *dbutils.DbRow
}

func (tuiApp *TuiApp) InitApp() {
	tuiApp.initPageSetup()
	tuiApp.createNewForm()
	tuiApp.selectInitPage()
	// tuiApp.displayEnteredForm()
	if err := tuiApp.App.SetRoot(tuiApp.Pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func (tuiApp *TuiApp) createNewForm() {

	tuiApp.NewTaskPage.
		AddTextArea("Task", tuiApp.Data.Task, 0, 3, 0, nil).
		AddTextArea("Description", tuiApp.Data.Description, 0, 15, 0, nil).
		AddCheckbox("Done", tuiApp.Data.Done, nil).
		AddCheckbox("High Priority", tuiApp.Data.HighPriority, nil).
		AddButton("Okay", func() {
			// tuiApp.saveData()
			tuiApp.getFormData()
			tuiApp.displayEnteredForm()
			// fmt.Println("Test - doing nothing for now...")
		}).
		AddButton("Menu", func() {
			tuiApp.Data = &dbutils.DbRow{}
			tuiApp.Pages.SwitchToPage("InitPage")
		}).
		AddButton("Quit", func() {
			tuiApp.App.Stop()
		}).
		AddButton("Clear", func() {
			tuiApp.clearFormData()
		})
	tuiApp.NewTaskPage.SetBorder(true).SetTitle("New Task").SetTitleAlign(tview.AlignLeft)

	tuiApp.Pages.AddPage("Create New Task", tuiApp.NewTaskPage, true, false)
}

func (tuiApp *TuiApp) initPageSetup() {
	tuiApp.Pages.AddPage("InitPage", tview.NewModal().
		SetText("Welcome to Todos").
		AddButtons([]string{"New Task", "Select Existing", "Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			switch buttonIndex {
			case 0:
				tuiApp.Data = &dbutils.DbRow{}
				tuiApp.Pages.SwitchToPage("Create New Task")
			case 1:
				tuiApp.Pages.SwitchToPage("SelectInitPage")
			default:
				tuiApp.App.Stop()
			}
		}), true, true)
}

func (tuiApp *TuiApp) displayEnteredForm() {
	newDisplayPage := tview.NewForm()
	newDisplayPage.
		AddTextView("Task", tuiApp.Data.Task, 0, 3, true, true).
		AddTextView("Description", tuiApp.Data.Description, 0, 15, true, true).
		AddTextView("Done", fmt.Sprintf("%t", tuiApp.Data.Done), 0, 0, true, false).
		AddTextView("High Priority", fmt.Sprintf("%t", tuiApp.Data.HighPriority), 0, 0, true, false).
		AddButton("Save", func() {
			tuiApp.Data.Description = time.Now().Local().String() + "\n" + tuiApp.Data.Description + "\n---"
			err := dbutils.InsertData(tuiApp.Sql, *tuiApp.Data)
			if err != nil {
				tuiApp.errorPage(err)
			}
			tuiApp.Data = &dbutils.DbRow{}
			tuiApp.Pages.SwitchToPage("InitPage")
		}).
		AddButton("Modify", func() {
			tuiApp.Pages.SwitchToPage("Create New Task")
		}).AddButton("Quit", func() {
		tuiApp.App.Stop()
	})
	newDisplayPage.SetFocus(4)
	newDisplayPage.SetBorder(true).SetTitle("Confirm Task").SetTitleAlign(tview.AlignLeft)
	tuiApp.Pages.AddAndSwitchToPage("Display Form", newDisplayPage, true)
}

func (tuiApp *TuiApp) getFormData() {
	taskTextArea := tuiApp.NewTaskPage.GetFormItem(0).(*tview.TextArea)
	descriptionTextArea := tuiApp.NewTaskPage.GetFormItem(1).(*tview.TextArea)
	doneCheckbox := tuiApp.NewTaskPage.GetFormItem(2).(*tview.Checkbox)
	priorityCheckbox := tuiApp.NewTaskPage.GetFormItem(3).(*tview.Checkbox)
	data := dbutils.NewInsertRow(
		taskTextArea.GetText(),
		descriptionTextArea.GetText(),
		doneCheckbox.IsChecked(),
		priorityCheckbox.IsChecked())

	tuiApp.Data = &data
}

func (tuiApp *TuiApp) clearFormData() {
	taskTextArea := tuiApp.NewTaskPage.GetFormItem(0).(*tview.TextArea)
	descriptionTextArea := tuiApp.NewTaskPage.GetFormItem(1).(*tview.TextArea)

	taskTextArea.SetText("", true)
	descriptionTextArea.SetText("", false)

	tuiApp.Data = &dbutils.DbRow{}
}

func (tuiApp *TuiApp) selectInitPage() {
	tuiApp.Pages.AddPage("SelectInitPage", tview.NewModal().
		SetText("Select Todos").
		AddButtons([]string{"Todo", "Done", "High Priority", "All", "Main Menu", "Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			switch buttonIndex {
			case 0:
				tuiApp.createSelectList("TodoTasks")
			case 1:
				tuiApp.createSelectList("DoneTasks")
			case 2:
				tuiApp.createSelectList("HighPriorityTasks")
			case 3:
				tuiApp.createSelectList("AllTasks")
			case 4:
				tuiApp.Data = &dbutils.DbRow{}
				tuiApp.Pages.SwitchToPage("InitPage")
			default:
				tuiApp.App.Stop()
			}
		}), false, false)
}

func (tuiApp *TuiApp) createSelectList(pageName string) {
	list := tview.NewList()
	databaseRows, err := dbutils.QuerySelector(tuiApp.Sql, pageName)
	if err != nil {
		tuiApp.errorPage(err)
	}
	run := 'a'
	for _, row := range databaseRows {
		list.
			AddItem(row.Task, row.Description, run, func() {
				tuiApp.selectForm(row)
			})
		run = run + 1
	}
	list.AddItem("Back to Menu", "Press for selection menu", '0', func() {
		tuiApp.Pages.SwitchToPage("SelectInitPage")
	})
	list.AddItem("Quit", "Press to exit", '1', func() {
		tuiApp.App.Stop()
	})
	tuiApp.Pages.AddAndSwitchToPage(pageName, list, true)
}

func (tuiApp *TuiApp) errorPage(err error) {
	tuiApp.Pages.AddAndSwitchToPage("Error Page", tview.NewModal().SetText(err.Error()).AddButtons([]string{"Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			tuiApp.App.Stop()
		}), true)
}

func (tuiApp *TuiApp) selectForm(dbRow dbutils.DbRow) {
	selectForm := tview.NewForm()
	selectForm.
		AddTextView("Task", dbRow.Task, 0, 3, true, true).
		AddTextArea("Description", dbRow.Description, 0, 6, 0, nil).
		AddTextArea("Add Description", "", 0, 8, 0, nil).
		AddCheckbox("Done", dbRow.Done, func(checkbox bool) {
			dbRow.Done = checkbox
		}).
		AddCheckbox("High Priority", dbRow.HighPriority, func(checkbox bool) {
			dbRow.HighPriority = checkbox
		}).
		AddButton("Save", func() {
			descriptionText := selectForm.GetFormItem(2).(*tview.TextArea).GetText()
			descriptionText = dbRow.Description + "\n" + time.Now().Local().String() + "\n" + descriptionText + "\n---"
			dbRow.Description = descriptionText
			dbRow.Done = selectForm.GetFormItem(3).(*tview.Checkbox).IsChecked()
			dbRow.HighPriority = selectForm.GetFormItem(4).(*tview.Checkbox).IsChecked()
			err := dbutils.UpdateRow(tuiApp.Sql, dbRow)
			if err != nil {
				tuiApp.errorPage(err)
			}
			tuiApp.Data = &dbutils.DbRow{}
			tuiApp.Pages.SwitchToPage("InitPage")
		}).
		AddButton("Delete", func() {
			err := dbutils.DeleteData(tuiApp.Sql, dbRow)
			if err != nil {
				tuiApp.errorPage(err)
			}
			tuiApp.Data = &dbutils.DbRow{}
			tuiApp.Pages.SwitchToPage("InitPage")
		}).
		AddButton("Menu", func() {
			tuiApp.Data = &dbutils.DbRow{}
			tuiApp.Pages.SwitchToPage("InitPage")
		}).
		AddButton("Quit", func() {
			tuiApp.App.Stop()
		}).SetBorder(true).SetTitle(fmt.Sprintf("Task %d", dbRow.Id)).SetTitleAlign(tview.AlignLeft)
	selectForm.SetFocus(1)
	tuiApp.Pages.AddAndSwitchToPage("SelectDataPage", selectForm, true)
}

// List Pages for the tasks - Four different pages - AddAndSwitchToPage
// - need to create this after the option is selected.

// After a task is selected - move to a form page for displaying the task and updation

// Update page for the task
