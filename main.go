package main

import (
	"type_tracker/utils"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/gdk"
	// "encoding/json"
	"log"
	// "net/http"
	// "net/url"
	"fmt"
	// "time"
	// "strconv"
)

const (
	Width = 800
	Height = 600

	Escape = 65307
	Shift = 65505
	Backspace = 65288
	Space = 32
)

type track struct {
	word string
	index uint
	length uint
	mistakesCount uint
	complete bool

	wordCount uint
	correctCount uint
	keystrokeCount uint
	badKeystrokeCount uint
}

var words *[]string
var tracker track
var keyMap = make(map[uint]uint)

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func initializeMap() {
	keyMap[Escape] = 0
	keyMap[Shift] = 1
	keyMap[Backspace] = 2
	keyMap[Space] = 3
}

func space() {
	if tracker.complete {
		tracker.correctCount++
	}
	tracker.wordCount++
	tracker.word = (*words)[tracker.wordCount]
	fmt.Printf("new word: %s\n", tracker.word)
	tracker.index = 0
	tracker.length = uint(len(tracker.word))
	tracker.mistakesCount = 0
	tracker.complete = false
}

func backspace() {
	if tracker.mistakesCount > 0 {
		tracker.mistakesCount--
	} else {
		tracker.badKeystrokeCount++
		if tracker.index > 0 {
			tracker.index--
		}
	}
}

func handleKeystroke(letter string, special bool, specialChar uint) {
	tracker.keystrokeCount++
	if special {
		switch specialChar {
		case 0:
			fmt.Println("Escape")
			//Reset
		case 1:
			fmt.Println("Shift")
			//Ignore this basically?
		case 2:
			fmt.Println("Backspace")
			backspace()
		case 3:
			fmt.Println("Space")
			fmt.Printf("mistakes: %d\nletters: %s\n", tracker.mistakesCount, string(tracker.word[tracker.index - 1]))
			space()
		}
	} else {
		
		if (letter == string(tracker.word[tracker.index]) && tracker.mistakesCount == 0) {
			fmt.Println("correct")
			tracker.index++
			if tracker.index == tracker.length {
				fmt.Println("correct word")
				tracker.complete = true
			} 
		} else {
			tracker.mistakesCount++
		}
	}
	fmt.Println("mistakes:", tracker.mistakesCount)

}

func keyEvent(entry *gtk.Entry, event *gdk.Event) {
	eventKey := gdk.EventKeyNewFromEvent(event)
	keyVal := eventKey.KeyVal()
	fmt.Println("keyVaL: ", keyVal)
	if val, ok := keyMap[keyVal]; ok {
		fmt.Println("special char")
		handleKeystroke("", true, val)
	}

	letter := string(keyVal)
	fmt.Println(letter)
	handleKeystroke(letter, false, 0)
}

func start() {
	gtk.Init(nil)
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	logError(err)
	
	win.SetTitle("type_tracker")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.SetDefaultSize(Width, Height)

	grid, err := gtk.GridNew()
	logError(err)

	var str string
	for i, w := range (*words)[:10] {
		if i == 5 {
			str += "\n"
		}
		w += " "
		str += w
	}

	frame1, err := gtk.FrameNew("words")
	logError(err)
	// frame1.SetShadowType(1)

	label1, err := gtk.LabelNew(str)
	logError(err)
	label1.SetName("label1")

	label1.SetVExpand(true)
	label1.SetHExpand(true)
	frame1.Add(label1)

	entry1, err := gtk.EntryNew()
	logError(err)
	entry1.SetVExpand(true)
	// entry1.SetHExpand(true)
	grid.SetRowSpacing(1)
	grid.SetColumnSpacing(1)
	grid.Attach(frame1, 0, 0, 400, 200)
	grid.Attach(entry1, 100, 200, 200, 100)

	//1 << 10
	entry1.AddEvents(1024)
	entry1.Connect("key_press_event", keyEvent)


	// style_provider, err := gtk.CssProviderNew()
	// logError(err)
	// style_provider.LoadFromPath("resources/style.css")
	// style_context, err := frame1.GetStyleContext()
	// logError(err)
	// // style_context.AddProvider(style_provider, 0)
	// screen, err := style_context.GetScreen()
	// logError(err)
	// gtk.AddProviderForScreen(screen, style_provider, 0)

	win.Add(grid)
	win.ShowAll()
	gtk.Main()
}


func main() {
	words = utils.GetNWords(300)

	tracker = track{word: (*words)[0], length: uint(len((*words)[0]))}
	initializeMap()
	// fmt.Println(*words)
	start()
}