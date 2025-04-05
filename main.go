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

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func updateFirstArrow(line *canvas.Line) {
	s := time.Now().Second()
	p := fyne.NewPos(float32(s), 0)
	line.Move(p)
}

func updateSecondArrow(line *canvas.Line) {
	s := (time.Now().Second() * 255) / 60
	//p := fyne.NewPos(0, float32(s))
	size := fyne.NewSize(255, float32(s))
	line.Resize(size)
	//line.Move(p)
}

func main() {
	var fc FClock
	a := app.New()
	fc.App = a
	fc.MainWindow = a.NewWindow("Fyne Clock")

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
	fc.MainContainer.Add(lineBegin)

	// The Second arrow of the pie
	lineEnd := canvas.NewLine(color.RGBA{22, 222, 12, 255})
	lineEnd.StrokeWidth = 5
	fc.MainContainer.Add(lineEnd)

	grid := container.New(layout.NewGridLayout(2), text1, text2, clockLabel, widget.NewLabel(""), c)

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
