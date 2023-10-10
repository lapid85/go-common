package utils

// InSlice 字符串序列是否包含
func InSlice[T int8 | uint8 | int16 | uint16 | int32 | uint32 | int | uint | int64 | uint64 | float32 | float64 | string | bool](arr []T, ele T) bool {
	for _, v := range arr {
		if ele == v {
			return true
		}
	}
	return false
}
