package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
)

var myapp fyne.App = app.New()

var myWindow fyne.Window = myapp.NewWindow("Devansh OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
// var img fyne.CanvasObject
// var DeskBtn fyne.Widget

var panelContent *fyne.Container

func main(){
	img := canvas.NewImageFromFile("E:\\Go-OS-main\\pic.jpg")
	myWindow.Resize(fyne.NewSize(1366, 768))

	weatherLogo, _ := fyne.LoadResourceFromPath("E:\\Go-OS-main\\logos\\weather.png")
	btn1 = widget.NewButtonWithIcon("", weatherLogo, func ()  {
		showWeatherApp()
	})

	// btn1.Resize(fyne.NewSize(10240, 10240))


	calLogo, _ := fyne.LoadResourceFromPath("E:\\Go-OS-main\\logos\\calc.png")
	btn2 = widget.NewButtonWithIcon("", calLogo, func ()  {
		showCalc()
	})

	// btn2.Resize(fyne.NewSize(600, 1600))


	galLogo, _ := fyne.LoadResourceFromPath("E:\\Go-OS-main\\logos\\gallery.png")
	btn3 = widget.NewButtonWithIcon("", galLogo, func ()  {
		showGalleryApp()
	})
	// btn3.Resize(fyne.NewSize(1550, 1250))


	textLogo, _ := fyne.LoadResourceFromPath("E:\\Go-OS-main\\logos\\texted.png")
	btn4 = widget.NewButtonWithIcon("", textLogo, func ()  {
		showTextEditor()
	})
	// btn4.Resize(fyne.NewSize(1050, 750))

	panelContent = container.NewVBox(layout.NewSpacer(), container.NewHBox(layout.NewSpacer(), btn1, btn2, btn3, btn4, layout.NewSpacer()), layout.NewSpacer(), layout.NewSpacer())
	myWindow.SetPadded(true)
	myWindow.CenterOnScreen()
	myWindow.SetContent(
		container.New(layout.NewMaxLayout(), img,
			container.NewBorder(nil, panelContent, nil, nil)),
	)
	myWindow.ShowAndRun()
}