package fun

// ClearSliceInt 去除slice的0值
func ClearSliceInt(slice []int,clearVal int) []int {
	var newSlice = make([]int, 0)
	for _, v := range slice {
		if v !=  clearVal {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// ClearSliceString 去除slice的""值
func ClearSliceString(slice []string,clearVal string) []string {
	var newSlice = make([]string, 0)
	for _, v := range slice {
		if v != clearVal {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
func InSliceInt(value int, slice []int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
func InSliceUint(value uint, slice []uint) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
func InSliceString(value string, slice []string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
func InSliceAny(value any, slice []any) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}


// SliceUniqueString String切片去重
func SliceUniqueString(value []string) []string {
	result := make([]string, 0, len(value))
	temp := map[string]struct{}{}
	for _, item := range value {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// SliceUniqueInt Int切片去重
func SliceUniqueInt(value []int) []int {
	result := make([]int, 0, len(value))
	temp := map[int]struct{}{}
	for _, item := range value {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// SliceUniqueUint Uint切片去重
func SliceUniqueUint(value []uint) []uint {
	result := make([]uint, 0, len(value))
	temp := map[uint]struct{}{}
	for _, item := range value {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
