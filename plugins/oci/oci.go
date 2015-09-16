package oci

import (
	"encoding/json"
	"fmt"
	"github.com/phoenix-io/phoenix-mon/plugins"
	ps "github.com/shirou/gopsutil/process"
	"io/ioutil"
)

type OCI struct {
	process []plugin.Process
}

type PState struct {
	ID  string `json:"id"`
	Pid int    `json:"init_process_pid"`
}

func init() {
	plugin.Register("oci", &plugin.RegisteredPlugin{New: NewPlugin})
}

func NewPlugin(pluginName string) (plugin.Plugin, error) {
	return &OCI{}, nil
}

func (p OCI) GetProcessList() ([]plugin.Process, error) {

	var state PState
	containers, _ := ioutil.ReadDir("/run/oci/")
	for _, c := range containers {
		//		fmt.Println(c.Name())
		data, err := ioutil.ReadFile("/run/oci/" + c.Name() + "/state.json")
		if err != nil {
			return p.process, fmt.Errorf("Error in reading file")
		}

		json.Unmarshal(data, &state)
		p.process = append(p.process, plugin.Process{Pid: state.Pid, Name: state.ID})
	}

	return p.process, nil
}

func (p OCI) GetProcessStat(process plugin.Process) (plugin.Process, error) {
	res, err := ps.NewProcess(int32(process.Pid))
	if err != nil {
		fmt.Println("Unable to Create Process Object")
	}
	memInfo, err := res.MemoryInfo()
	if err != nil {
		return process, fmt.Errorf("geting ppid error %v", err)

	}
	process.Memory = memInfo.RSS
	return process, nil
}
