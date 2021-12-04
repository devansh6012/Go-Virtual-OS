package main

import (
	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/theme"
)

func showGalleryApp() {
	// a := app.New()
	w := myapp.NewWindow("Hello")
	w.Resize(fyne.NewSize(1280,720))
	root_src:="C:\\Users\\compass1415\\Pictures"
	
	files, err := ioutil.ReadDir(root_src)
    if err != nil {
        log.Fatal(err)
    }

	tabs := container.NewAppTabs()
	for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir(){
			extension:= strings.Split(file.Name(),".")[1]
			if extension == "png" || extension == "jpg" || extension == "jpeg"{
				tabs.Append(container.NewTabItem(file.Name(),canvas.NewImageFromFile(root_src+"\\"+file.Name())))
			}
		}
    }	

	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs);
	w.Show()
}