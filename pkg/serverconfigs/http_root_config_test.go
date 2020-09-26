package serverconfigs

import (
	"github.com/iwind/TeaGo/assert"
	"testing"
)

func TestHTTPRootConfig_HasVariables(t *testing.T) {
	a := assert.NewAssertion(t)

	{
		rootConfig := &HTTPRootConfig{
			Dir: "",
		}
		err := rootConfig.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsFalse(rootConfig.HasVariables())
	}

	{
		rootConfig := &HTTPRootConfig{
			Dir: "/home/www",
		}
		err := rootConfig.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsFalse(rootConfig.HasVariables())
	}

	{
		rootConfig := &HTTPRootConfig{
			Dir: "/home/www/${prefix}/world",
		}
		err := rootConfig.Init()
		if err != nil {
			t.Fatal(err)
		}
		a.IsTrue(rootConfig.HasVariables())
	}
}
