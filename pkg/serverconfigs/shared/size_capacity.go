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
		return this.Count * 1024
	case SizeCapacityUnitMB:
		return this.Count * 1024 * 1024
	case SizeCapacityUnitGB:
		return this.Count * 1024 * 1024 * 1024
	case SizeCapacityUnitTB:
		return this.Count * 1024 * 1024 * 1024 * 1024
	case SizeCapacityUnitPB:
		return this.Count * 1024 * 1024 * 1024 * 1024 * 1024
	default:
		return this.Count
	}
}

func (this *SizeCapacity) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}
