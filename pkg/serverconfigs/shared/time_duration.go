package shared

import (
	"encoding/json"
	"github.com/iwind/TeaGo/types"
	"time"
)

type TimeDurationUnit = string

const (
	TimeDurationUnitMS     TimeDurationUnit = "ms"
	TimeDurationUnitSecond TimeDurationUnit = "second"
	TimeDurationUnitMinute TimeDurationUnit = "minute"
	TimeDurationUnitHour   TimeDurationUnit = "hour"
	TimeDurationUnitDay    TimeDurationUnit = "day"
	TimeDurationUnitWeek   TimeDurationUnit = "week"
)

// TimeDuration 时间间隔
type TimeDuration struct {
	Count int64            `yaml:"count" json:"count"` // 数量
	Unit  TimeDurationUnit `yaml:"unit" json:"unit"`   // 单位
}

func (this *TimeDuration) Duration() time.Duration {
	switch this.Unit {
	case TimeDurationUnitMS:
		return time.Duration(this.Count) * time.Millisecond
	case TimeDurationUnitSecond:
		return time.Duration(this.Count) * time.Second
	case TimeDurationUnitMinute:
		return time.Duration(this.Count) * time.Minute
	case TimeDurationUnitHour:
		return time.Duration(this.Count) * time.Hour
	case TimeDurationUnitDay:
		return time.Duration(this.Count) * 24 * time.Hour
	case TimeDurationUnitWeek:
		return time.Duration(this.Count) * 24 * 7 * time.Hour
	default:
		return time.Duration(this.Count) * time.Second
	}
}

func (this *TimeDuration) Seconds() int64 {
	switch this.Unit {
	case TimeDurationUnitMS:
		return this.Count / 1000
	case TimeDurationUnitSecond:
		return this.Count
	case TimeDurationUnitMinute:
		return this.Count * 60
	case TimeDurationUnitHour:
		return this.Count * 3600
	case TimeDurationUnitDay:
		return this.Count * 3600 * 24
	case TimeDurationUnitWeek:
		return this.Count * 3600 * 24 * 7
	default:
		return this.Count
	}
}

func (this *TimeDuration) Description() string {
	var countString = types.String(this.Count)
	switch this.Unit {
	case TimeDurationUnitMS:
		return countString + "毫秒"
	case TimeDurationUnitSecond:
		return countString + "秒"
	case TimeDurationUnitMinute:
		return countString + "分钟"
	case TimeDurationUnitHour:
		return countString + "小时"
	case TimeDurationUnitDay:
		return countString + "天"
	case TimeDurationUnitWeek:
		return countString + "周"
	default:
		return countString + "秒"
	}
}

func (this *TimeDuration) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}
