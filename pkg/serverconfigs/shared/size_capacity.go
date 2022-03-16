package shared

import "encoding/json"

type SizeCapacityUnit = string

const (
	SizeCapacityUnitByte SizeCapacityUnit = "byte"
	SizeCapacityUnitKB   SizeCapacityUnit = "kb"
	SizeCapacityUnitMB   SizeCapacityUnit = "mb"
	SizeCapacityUnitGB   SizeCapacityUnit = "gb"
	SizeCapacityUnitTB   SizeCapacityUnit = "tb"
	SizeCapacityUnitPB   SizeCapacityUnit = "pb"
	SizeCapacityUnitEB   SizeCapacityUnit = "eb"
	//SizeCapacityUnitZB   SizeCapacityUnit = "zb" // zb和yb超出int64范围，暂不支持
	//SizeCapacityUnitYB   SizeCapacityUnit = "yb"
)

type SizeCapacity struct {
	Count int64            `json:"count" yaml:"count"`
	Unit  SizeCapacityUnit `json:"unit" yaml:"unit"`
}

func (this *SizeCapacity) Bytes() int64 {
	if this.Count < 0 {
		return -1
	}
	switch this.Unit {
	case SizeCapacityUnitByte:
		return this.Count
	case SizeCapacityUnitKB:
		return this.Count * this.pow(1)
	case SizeCapacityUnitMB:
		return this.Count * this.pow(2)
	case SizeCapacityUnitGB:
		return this.Count * this.pow(3)
	case SizeCapacityUnitTB:
		return this.Count * this.pow(4)
	case SizeCapacityUnitPB:
		return this.Count * this.pow(5)
	case SizeCapacityUnitEB:
		return this.Count * this.pow(6)
	default:
		return this.Count
	}
}

func (this *SizeCapacity) IsNotEmpty() bool {
	return this.Count > 0
}

func (this *SizeCapacity) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}

func (this *SizeCapacity) pow(n int) int64 {
	if n <= 0 {
		return 1
	}
	if n == 1 {
		return 1024
	}
	return this.pow(n-1) * 1024
}
