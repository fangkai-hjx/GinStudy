package validCheck

import "strconv"

type StrTo string

// 转化为 String
func (s StrTo) String() string {
	return string(s)
}

// 转化为 UInt32
func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

// 转化为 UInt64
func (s StrTo) UInt64() (uint64, error) {
	v, err := strconv.Atoi(s.String())
	return uint64(v), err
}

// 转化为 Int
func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}

// 验证UInt64
func (s StrTo) MustUInt64() uint64 {
	v, _ := s.UInt64()
	return v
}
