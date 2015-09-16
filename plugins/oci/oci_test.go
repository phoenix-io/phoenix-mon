package oci

import "testing"
import _ "github.com/phoenix-io/phoenix-mon/plugins"

func TestGetProcessList(t *testing.T) {
	var p OCI
	_, err := p.GetProcessList()
	if err != nil {
		t.Errorf("Unable to fetch the processlist")
	}
}

func TestGetProcessStat(t *testing.T) {
	var p OCI
	plist, err := p.GetProcessList()
	if err != nil {
		t.Errorf("Unable to fetch the processlist")
	}

	for _, c := range plist {
		_, err = p.GetProcessStat(c)
		if err != nil {
			t.Errorf("Unable to fetech process info")
		}
	}
}
