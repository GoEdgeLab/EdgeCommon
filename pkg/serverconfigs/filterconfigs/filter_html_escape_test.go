package filterconfigs

import "testing"

func TestHTMLEscapeFilter_Do(t *testing.T) {
	filter := &HTMLEscapeFilter{}
	t.Log(filter.Do("Hello", nil))
	t.Log(filter.Do("<script></script>", nil))
}
