package filterconfigs

import "testing"

func TestToBytes(t *testing.T) {
	t.Log(ToBytes("hello"))
	t.Log(ToBytes(123))
	t.Log(ToBytes([]byte{1, 2, 3}))
}
