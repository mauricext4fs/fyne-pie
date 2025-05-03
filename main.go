package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type FClock struct {
	App           fyne.App
	MainWindow    fyne.Window
	MainContainer *fyne.Container
}

var fc FClock

func main() {
	a := app.New()
	fc.App = a
	fc.MainWindow = fc.App.NewWindow("Fyne Pie")

	clockLabel := widget.NewLabel("")
	updateTime(clockLabel)

	text1 := canvas.NewText("Vlad", color.Black)
	text2 := canvas.NewText("Tepes", color.Black)

	c := container.NewStack()
	fc.MainContainer = c

	// The ring
	circle := canvas.NewCircle(color.Transparent)
	circle.StrokeColor = color.Gray{Y: 0x99}
	circle.StrokeWidth = 15
	fc.MainContainer.Add(circle)

	// The First arrow of the pie
	lineBegin := canvas.NewLine(color.RGBA{255, 22, 12, 255})
	lineBegin.StrokeWidth = 50
	//fc.MainContainer.Add(lineBegin)

	// The Second arrow of the pie
	lineEnd := canvas.NewLine(color.RGBA{22, 222, 12, 255})
	lineEnd.StrokeWidth = 5
	//fc.MainContainer.Add(lineEnd)

	// Clocki
	clockC := drawSomeClock()
	clockC.Layout = fc.MainContainer.Layout
	fc.MainContainer.Add(clockC)

	grid := container.New(layout.NewGridLayout(2), text1, text2, clockLabel, widget.NewLabel(""), fc.MainContainer)

	fc.MainWindow.SetContent(grid)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clockLabel)
			updateFirstArrow(lineBegin)
			updateSecondArrow(lineEnd)
		}
	}()

	fc.MainWindow.Resize(fyne.NewSize(700, 700))
	fc.MainWindow.ShowAndRun()

}
