package main

import (
	"type_tracker/utils"
	"github.com/gotk3/gotk3/gtk"
	// "encoding/json"
	"log"
	// "net/http"
	// "net/url"
	"fmt"
)

func start() {
	fmt.Println("HERE")
	gtk.Init(nil)
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal(err)
	}
	
	win.SetTitle("type_tracker")
	win.Connewct("destroy", func() {
		gtk.MainQuit()
	})

	win.SetDefaultSize(800, 600)
	win.showAll()
	gtk.Main()
}


func main() {
	words := utils.GetNWords(300)
	start()
}