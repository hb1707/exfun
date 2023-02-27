package fun

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Round 小数点四舍五入
func Round(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}

//IntChecked 值是否存在
func IntChecked(selectInt int, sumInt int) bool {
	return selectInt&sumInt == selectInt
}

//RandInt 生成指定位的随机整数
func RandInt(n int64) int64 {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	result, _ := strconv.ParseInt(string(b), 10, 64)
	return result
}
