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

//RandInt 生成指定位数的随机数
func RandInt(n int) int {
	return int(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(10 * n)))
}
