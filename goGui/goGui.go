package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

func main() {
	//create app & add new window
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	//add tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	//show tabs
	myWindow.SetContent(tabs)

	//resize window
	myWindow.Resize(fyne.NewSize(500, 500))

	//run app
	myWindow.ShowAndRun()
}
