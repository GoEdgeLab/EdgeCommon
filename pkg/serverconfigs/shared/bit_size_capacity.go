package shared

import "encoding/json"

type BitSizeCapacityUnit = string

const (
	BitSizeCapacityUnitB  BitSizeCapacityUnit = "b"
	BitSizeCapacityUnitKB BitSizeCapacityUnit = "kb"
	BitSizeCapacityUnitMB BitSizeCapacityUnit = "mb"
	BitSizeCapacityUnitGB BitSizeCapacityUnit = "gb"
	BitSizeCapacityUnitTB BitSizeCapacityUnit = "tb"
	BitSizeCapacityUnitPB BitSizeCapacityUnit = "pb"
	BitSizeCapacityUnitEB BitSizeCapacityUnit = "eb"
	//BitSizeCapacityUnitZB   BitSizeCapacityUnit = "zb" // zb和yb超出int64范围，暂不支持
	//BitSizeCapacityUnitYB   BitSizeCapacityUnit = "yb"
)

type BitSizeCapacity struct {
	Count int64               `json:"count" yaml:"count"`
	Unit  BitSizeCapacityUnit `json:"unit" yaml:"unit"`
}

func NewBitSizeCapacity(count int64, unit BitSizeCapacityUnit) *BitSizeCapacity {
	return &BitSizeCapacity{
		Count: count,
		Unit:  unit,
	}
}

func DecodeBitSizeCapacityJSON(sizeCapacityJSON []byte) (*BitSizeCapacity, error) {
	var capacity = &BitSizeCapacity{}
	err := json.Unmarshal(sizeCapacityJSON, capacity)
	return capacity, err
}

func (this *BitSizeCapacity) Bits() int64 {
	if this.Count < 0 {
		return -1
	}
	switch this.Unit {
	case BitSizeCapacityUnitB:
		return this.Count
	case BitSizeCapacityUnitKB:
		return this.Count * this.pow(1)
	case BitSizeCapacityUnitMB:
		return this.Count * this.pow(2)
	case BitSizeCapacityUnitGB:
		return this.Count * this.pow(3)
	case BitSizeCapacityUnitTB:
		return this.Count * this.pow(4)
	case BitSizeCapacityUnitPB:
		return this.Count * this.pow(5)
	case BitSizeCapacityUnitEB:
		return this.Count * this.pow(6)
	default:
		return this.Count
	}
}

func (this *BitSizeCapacity) IsNotEmpty() bool {
	return this.Count > 0
}

func (this *BitSizeCapacity) AsJSON() ([]byte, error) {
	return json.Marshal(this)
}

func (this *BitSizeCapacity) pow(n int) int64 {
	if n <= 0 {
		return 1
	}
	if n == 1 {
		return 1024 // TODO 考虑是否使用1000进制
	}
	return this.pow(n-1) * 1024
}
