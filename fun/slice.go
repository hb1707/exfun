package fun

// ClearZeroSliceInt 去除slice的0值
func ClearZeroSliceInt(slice []int) []int {
	var newSlice = make([]int, 0, 0)
	for _, v := range slice {
		if v > 0 {
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

// SliceUniqueString String切片去重
func SliceUniqueString(value *[]string) *[]string {
	result := make([]string, 0, len(*value))
	temp := map[string]struct{}{}
	for _, item := range *value {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return &result
}

// SliceUniqueInt Int切片去重
func SliceUniqueInt(value *[]int) *[]int {
	result := make([]int, 0, len(*value))
	temp := map[int]struct{}{}
	for _, item := range *value {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return &result
}
