package nodeconfigs

import "testing"

func TestNodeStatus_ComputerBuildVersionCode(t *testing.T) {
	{
		status := &NodeStatus{}
		status.ComputerBuildVersionCode()
		t.Log(status.BuildVersion, status.BuildVersionCode)
	}

	{
		status := &NodeStatus{BuildVersion: "0.0.6"}
		status.ComputerBuildVersionCode()
		t.Log(status.BuildVersion, status.BuildVersionCode)
	}

	{
		status := &NodeStatus{BuildVersion: "0.0.6.1"}
		status.ComputerBuildVersionCode()
		t.Log(status.BuildVersion, status.BuildVersionCode)
	}

	{
		status := &NodeStatus{BuildVersion: "0.0.7"}
		status.ComputerBuildVersionCode()
		t.Log(status.BuildVersion, status.BuildVersionCode)
	}

	{
		status := &NodeStatus{BuildVersion: "0.7"}
		status.ComputerBuildVersionCode()
		t.Log(status.BuildVersion, status.BuildVersionCode)
	}
	{
		status := &NodeStatus{BuildVersion: "7"}
		status.ComputerBuildVersionCode()
		t.Log(status.BuildVersion, status.BuildVersionCode)
	}
	{
		status := &NodeStatus{BuildVersion: "7.0.1"}
		status.ComputerBuildVersionCode()
		t.Log(status.BuildVersion, status.BuildVersionCode)
	}
}
