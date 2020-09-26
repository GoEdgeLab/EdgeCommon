package serverconfigs

import (
	"regexp"
	"testing"
	"time"
)

func TestHTTPLocationConfig_Match_Reg(t *testing.T) {
	randString := "1024"
	reg := regexp.MustCompile(`(?P<num>\d+)`)
	before := time.Now()
	subNames := reg.SubexpNames()
	match := reg.FindStringSubmatch(randString)
	t.Log(time.Since(before).Seconds()*1000, "ms")
	t.Log(subNames[1], "=", match[1])
}
