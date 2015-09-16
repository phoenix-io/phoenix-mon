package main

import (
	_ "fmt"
	"time"

	ui "github.com/gizak/termui"
	box "github.com/nsf/termbox-go"
	"github.com/phoenix-io/phoenix-mon/plugins"
	_ "github.com/phoenix-io/phoenix-mon/plugins/oci"
)


func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	ui.UseTheme("helloworld")

	header := createHeader("-- Phoenix Dashboard --")
	footer := createFooterBar()

	plugin, _ := plugin.NewPlugin("oci")
	plugin.GetProcessList()

	paintScreen := func() {
		ui.Render(header, footer)
	}

	evt := ui.EventCh()
	for {
		select {
		case e := <-evt:
			if e.Ch == 'q' {
				return
			}
		default:
			box.Sync()
			paintScreen()
			time.Sleep(time.Second)
		}
	}
}

func createHeader(msg string) *ui.Par {

	w, _ := box.Size()
	p := ui.NewPar(msg)
	p.HasBorder = false
	p.Height = 3
	p.Width = len(msg)
	p.Y = 0
	p.X = w/2 - len(msg)/2
	return p
}

func createFooterBar() *ui.Par {
	_, h := box.Size()
	footer := ui.NewPar("press 'q' to quit")
	footer.Height = 3
	footer.Width = 20
	footer.Y = h - 3

	return footer
}
