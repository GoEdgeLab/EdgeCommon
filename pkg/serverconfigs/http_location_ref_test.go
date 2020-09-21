package serverconfigs

import (
	"github.com/iwind/TeaGo/logs"
	"testing"
)

func TestHTTPLocationRef(t *testing.T) {
	ref := &HTTPLocationRef{LocationId: 1}
	logs.PrintAsJSON(ref, t)

	ref.Children = append(ref.Children, &HTTPLocationRef{LocationId: 2}, &HTTPLocationRef{LocationId: 3})
	logs.PrintAsJSON(ref, t)
}
