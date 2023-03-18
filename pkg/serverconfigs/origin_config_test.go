package serverconfigs

import "testing"

func TestOriginConfig_UniqueKey(t *testing.T) {
	origin := &OriginConfig{
		Id:      1,
		Version: 101,
	}
	err := origin.Init(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(origin.UniqueKey())
}
