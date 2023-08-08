package serverconfigs

import (
	"context"
	"testing"
)

func TestOriginConfig_UniqueKey(t *testing.T) {
	origin := &OriginConfig{
		Id:      1,
		Version: 101,
	}
	err := origin.Init(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(origin.UniqueKey())
}
