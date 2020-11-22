package filterconfigs

import "testing"

func TestSha256Filter_Do(t *testing.T) {
	filter := &Sha256Filter{}
	t.Log(filter.Do("123456", nil))
	t.Log(filter.Do("", nil))
}
