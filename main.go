package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	app := app.New()                   //inits the app
	window := app.NewWindow("AlgoVis") //titles the window
	window.Resize(fyne.NewSize(1200, 800))

	content := initDllEnv()
	window.SetContent(content)
	window.ShowAndRun()
}

func initDllEnv() *fyne.Container {
	circle := canvas.NewCircle(color.White)
	circle.StrokeColor = color.White
	circle.StrokeWidth = 5
	circle.Resize(fyne.NewSize(100, 100))
	cirWid := circle.Size().Width
	cirHei := circle.Size().Height

	label := canvas.NewText("Node {}", color.Black)
	label.Alignment = fyne.TextAlignCenter
	labWid := label.MinSize().Width
	labHei := label.MinSize().Height

	container := container.NewWithoutLayout(circle, label)
	container.Resize(fyne.NewSize(cirWid, cirHei))

	label.Resize(label.MinSize())
	label.Move(fyne.NewPos((cirWid-labWid)/2, (cirHei-labHei)/2))
	return container
}
