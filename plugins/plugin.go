package plugin

import (
	"fmt"
)

type Plugin interface {
	GetProcessList() ([]Process, error)
	GetProcessStat(Process) (Process, error)
}

type RegisteredPlugin struct {
	New func(pluginName string) (Plugin, error)
}

type Process struct {
	Pid    int
	Name   string
	Cpu    uint
	Memory uint64
}

var (
	plugins map[string]*RegisteredPlugin
)

func init() {
	plugins = make(map[string]*RegisteredPlugin)
}

func Register(name string, registeredPlugin *RegisteredPlugin) error {

	if _, exists := plugins[name]; exists {
		return fmt.Errorf("Plugin already registered %s", name)
	}

	plugins[name] = registeredPlugin
	return nil
}

func NewPlugin(name string) (Plugin, error) {
	plugin, exists := plugins[name]
	if !exists {
		return nil, fmt.Errorf("Plugin: Unknown plugin %s", name)
	}

	return plugin.New(name)
}
