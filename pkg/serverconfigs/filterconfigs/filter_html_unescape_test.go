package filterconfigs

import "testing"

func TestHTMLUnescapeFilter_Do(t *testing.T) {
	filter := &HTMLUnescapeFilter{}
	t.Log(filter.Do("Hello", nil))
	t.Log(filter.Do("&lt;script&gt;", nil))
	t.Log(filter.Do("<script>", nil))
}
