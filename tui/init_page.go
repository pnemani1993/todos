package tui

import (
	"fmt"

	"github.com/rivo/tview"
)

func InitButtons() {
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("New Task", "To create a new task", 'n', nil).
		AddItem("Select Existing", "To select an existing task", 's', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func TryingPages() {
	app := tview.NewApplication()
	pages := tview.NewPages()
	for page := 0; page < 5; page++ {
		func(page int) {
			pages.AddPage(fmt.Sprintf("page-%d", page),
				tview.NewModal().
					SetText(fmt.Sprintf("This is page %d. Choose where to go next.", page+1)).
					AddButtons([]string{"Next", "Quit"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						if buttonIndex == 0 {
							pages.SwitchToPage(fmt.Sprintf("page-%d", (page+1)%5))
						} else {
							app.Stop()
						}
					}),
				false,
				page == 0)
		}(page)
	}
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

// InitPage
func GetInitModal(app *tview.Application, pages *tview.Pages) {

	pages.AddPage("InitPage", tview.NewModal().
		SetText("Welcome to Todos").
		AddButtons([]string{"New Task", "Select Existing", "Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonIndex == 0 {
				pages.SwitchToPage("Create New Task")
			} else if buttonIndex == 1 {
				pages.SwitchToPage("View Or Update")
			} else {
				app.Stop()
			}
		}), false, true)
}

// Create New Task - For creating a new task
func CreateNewTaskForm(app *tview.Application, pages tview.Pages) *tview.Form {
	form := tview.NewForm().
		// AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
		AddInputField("First name", "", 20, nil, nil).
		AddInputField("Last name", "", 20, nil, nil).
		AddTextArea("Task", "Enter task", 30, 0, 0, nil).
		AddTextArea("Description", "Enter Task description", 0, 0, 0, nil).
		// AddTextView("Notes", "This is just a demo.\nYou can enter whatever you wish.", 40, 2, true, false).
		AddCheckbox("Age 18+", false, nil).
		AddPasswordField("Password", "", 10, '*', nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
	return form
}
