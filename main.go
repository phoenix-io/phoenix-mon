package main

import (
	ui "github.com/gizak/termui"
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
	plist, _ := plugin.GetProcessList()

	pls := createProcessList(plist)

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 5, header)),
		ui.NewRow(
			ui.NewCol(6, 0, pls)),
		ui.NewRow(
			ui.NewCol(12, 0, footer)),
	)

	ui.Body.Align()
	ui.Render(ui.Body)

	paintScreen := func() {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Render(ui.Body)
	}

	evt := ui.EventCh()
	for {
		select {
		case e := <-evt:
			if e.Ch == 'q' {
				return
			}
			if e.Type == ui.EventResize {
				go paintScreen()
			}
		}
	}
}

func createProcessList(plist []plugin.Process) *ui.List {
	ls := ui.NewList()

	for _, p := range plist {
		ls.Items = append(ls.Items, p.Name)
	}
	ls.ItemFgColor = ui.ColorYellow
	ls.Border.Label = "Process List"
	ls.Height = 10
	return ls
}

func createHeader(msg string) *ui.Par {
	p := ui.NewPar(msg)
	p.HasBorder = false
	p.Height = 3
	return p
}

func createFooterBar() *ui.Par {
	footer := ui.NewPar("press 'q' to quit")
	footer.Height = 3

	return footer
}
