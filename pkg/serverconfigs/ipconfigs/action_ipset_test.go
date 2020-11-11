package ipconfigs

import "testing"

func TestIPSetAction_Run(t *testing.T) {
	action := &IPSetAction{}
	err := action.Run(&IPItemConfig{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
