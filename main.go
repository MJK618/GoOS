package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	myOS = app.New()
)
var myOSWindow = myOS.NewWindow("GoOS")

var tab1 fyne.Widget
var tab2 fyne.Widget
var tab3 fyne.Widget
var tab4 fyne.Widget

var homeBGImg fyne.CanvasObject
var homeButton fyne.Widget
var appPanel *fyne.Container

func main() {
	//myOS.Settings().SetTheme(theme.LightTheme())
	homeBGImg = canvas.NewImageFromFile("systemfiles/os/img/OSbg.jpg")
	tab1 = widget.NewButtonWithIcon("Weather", theme.InfoIcon(), func() {
		showWeatherApp(myOSWindow)
	})
	tab2 = widget.NewButtonWithIcon("Calculator", theme.FileApplicationIcon(), func() {
		showCalculatorApp()
	})
	tab3 = widget.NewButtonWithIcon("Gallery", theme.FileImageIcon(), func() {
		showGalleryApp(myOSWindow)
	})
	tab4 = widget.NewButtonWithIcon("GoPad", theme.DocumentIcon(), func() {
		showGoPadApp()
	})
	homeButton = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		myOSWindow.SetContent(container.NewBorder( nil,nil,appPanel,nil,homeBGImg))
	})

	appPanel = container.NewHBox(container.NewGridWithRows(5, homeButton, tab1,tab2,tab3,tab4))

	myOSWindow.Resize(fyne.NewSize(1200,700))
	myOSWindow.CenterOnScreen()
	myOSWindow.SetContent(container.NewBorder(nil,nil,appPanel,nil,homeBGImg))
	myOSWindow.ShowAndRun()
}
