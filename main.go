package main

import (
	"fmt"
	ui "github.com/gizak/termui"
	"github.com/phoenix-io/phoenix-mon/plugins"
	_ "github.com/phoenix-io/phoenix-mon/plugins/oci"
)

func main() {

	var memStat []float64
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	ui.UseTheme("helloworld")

	header := createHeader("-- Phoenix Dashboard --")
	footer := createFooterBar()

	p, _ := plugin.NewPlugin("oci")
	procList := processList(p)
	pls := createProcessList(procList)

	plist, _ := p.GetProcessList()
	for _, pl := range plist {

		memStat = append(memStat, processMemory(p, pl))
		break
	}
	memWidget := createProcessMemoryWidget(memStat)

	// Adjust widgets on screen
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 5, header)),
		ui.NewRow(
			ui.NewCol(3, 0, pls),
			ui.NewCol(4, 0, memWidget)),
		ui.NewRow(
			ui.NewCol(12, 0, footer)),
	)

	ui.Body.Align()
	ui.Render(ui.Body)

	paintScreen := func() {
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		pls.Items = processList(p)
		plist, _ := p.GetProcessList()
		for _, pl := range plist {

			memStat = append(memStat, processMemory(p, pl))
			break
		}
		memStat = []float64{1,2,3,4,5,6,7,8,9,8,7,6,4,3,5,7,9,4,3,2,4,3}
		memWidget.Data = memStat
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

func processList(p plugin.Plugin) []string {
	var processList []string
	plist, _ := p.GetProcessList()
	for _, p := range plist {
		name := fmt.Sprintf("%s", p.Name)
		processList = append(processList, name)
	}

	return processList
}

func processMemory(p plugin.Plugin, process plugin.Process) float64 {

	p.GetProcessStat(process)
	return 6
	//return float64(mem.Memory)
}
