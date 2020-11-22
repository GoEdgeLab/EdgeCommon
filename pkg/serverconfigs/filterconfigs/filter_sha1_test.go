package filterconfigs

import "testing"

func TestSha1Filter_Do(t *testing.T) {
	filter := &Sha1Filter{}
	t.Log(filter.Do("123456", nil))
	t.Log(filter.Do("", nil))
}
