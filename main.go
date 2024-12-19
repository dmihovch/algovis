package main

import (
	"algovis/dll"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()                   //inits the app
	window := app.NewWindow("AlgoVis") //titles the window
	window.Resize(fyne.NewSize(1200, 800))

	label := widget.NewLabel("Print DLL to console")
	//label.Move(fyne.NewPos(100, 50))
	//label.Resize(fyne.NewSize(200, 100))
	LinkedList := widget.NewButton("click me", func() { dll.DllInit() })
	//LinkedList.Move(fyne.NewPos(100, 100))
	//LinkedList.Resize(fyne.NewSize(200, 200))
	sideMenu := container.NewVBox(label, LinkedList)
	separator := widget.NewSeparator()
	contentSpace := container.NewCenter(widget.NewLabel("DLLs in action"), widget.NewButton("clicky button", func() { fmt.Println("anything") }))
	content := container.NewHBox(sideMenu, separator, contentSpace)

	window.SetContent(content)
	window.ShowAndRun()
}
