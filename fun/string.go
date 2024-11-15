package fun

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"strconv"
)
// Str2Int 字符串转换为整数
func Str2Int(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}


// ClearTags 清除代码
func ClearTags(str string) string {
	str = StripTags(str)
	//去除换行及其他特殊字符
	re, _ := regexp.Compile(`(\r|\n|\t)`)
	str = re.ReplaceAllString(str, "")
	str = strings.Trim(str, " ")
	return str
}

// IsChineseAndEnglish 判断是否只有中英文和_
func IsChineseAndEnglish(str string, dot bool) bool {
	if dot {
		matched, _ := regexp.MatchString("^[_0-9a-zA-Z\u4e00-\u9fa5.]+$", str)
		return matched
	}
	matched, _ := regexp.MatchString("^[_0-9a-zA-Z\u4e00-\u9fa5]+$", str)
	return matched
}

// IsEnglish 判断是否只有英文和_
func IsEnglish(str string, dot bool) bool {
	if dot {
		matched, _ := regexp.MatchString("^[_0-9a-zA-Z.]+$", str)
		return matched
	}
	matched, _ := regexp.MatchString("^[_0-9a-zA-Z]+$", str)
	return matched
}

// StripTags strip_tags
func StripTags(str string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile(`\<[\S\s]+?\>`)
	str = re.ReplaceAllStringFunc(str, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile(`\<style[\S\s]+?\</style\>`)
	str = re.ReplaceAllString(str, "")

	//去除SCRIPT
	re, _ = regexp.Compile(`\<script[\S\s]+?\</script\>`)
	str = re.ReplaceAllString(str, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile(`\<[\S\s]+?\>`)
	str = re.ReplaceAllString(str, "")
	return str
}

//substr 截取字符串，支持中文
func SubStr(str string, start int, length int) string {
	rs := []rune(str)
	max := len(rs)
	end := length
	if start < 0 {
		start = max + start //start负数 - 在从字符串结尾开始的指定位置开始
	}
	if length < 0 {
		end = max + length //length负数 - 从字符串末端返回的长度
	} else if length == 0 {
		end = max
	} else {
		end = start + length //start正数 - 从 start 参数所在的位置返回的长度
	}
	if end < 0 {
		end = 0
	}
	if start < 0 {
		start = 0
	}
	if start > max {
		start = max
	}
	if end > max {
		end = max
	}
	return string(rs[start:end])
}

// addslashes() 函数返回在预定义字符之前添加反斜杠的字符串。
// 预定义字符是：
// 单引号（'）
// 双引号（"）
// 反斜杠（\）
func Addslashes(str string) string {
	tmpRune := []rune{}
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
			tmpRune = append(tmpRune, ch)
		default:
			tmpRune = append(tmpRune, ch)
		}
	}
	return string(tmpRune)
}

// stripslashes() 函数删除由 addslashes() 函数添加的反斜杠。
func Stripslashes(str string) string {
	dstRune := []rune{}
	strRune := []rune(str)
	strLenth := len(strRune)
	for i := 0; i < strLenth; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}

//Stripos 查找字符串在另一字符串中第一次出现的位置（不区分大小写）
func Stripos(str string, index string) int {
	return strings.Index(strings.ToLower(str), strings.ToLower(index))
}

//Strpos 查找字符串在另一字符串中第一次出现的位置（区分大小写）
func Strpos(str string, index string) int {
	return strings.Index(str, index)
}

//Strripos 查找字符串在另一字符串中最后一次出现的位置（不区分大小写）
func Strripos(str string, index string) int {
	return strings.LastIndex(strings.ToLower(str), strings.ToLower(index))
}

//Strrpos 查找字符串在另一字符串中最后一次出现的位置（区分大小写）
func Strrpos(str string, index string) int {
	return strings.LastIndex(str, index)
}

func PregMatchAll(pattern string, subject string, matches *[][]string) bool {
	data := regexp.MustCompile(pattern).FindAllStringSubmatch(subject, -1)

	matchAll := make(map[int][]string, 2)
	for _, va := range data {
		for kb, vb := range va {
			matchAll[kb] = append(matchAll[kb], vb)
		}
	}
	for _, a := range matchAll {
		*matches = append(*matches, a)
	}

	if *matches == nil {
		return false
	} else {
		return true
	}
}

func PregMatch(pattern string, subject string, matches *[]string) bool {
	if matches == nil {
		return false
	}
	data := regexp.MustCompile(pattern).FindStringSubmatch(subject)
	for _, match := range data {
		*matches = append(*matches, match)
	}

	if *matches == nil {
		return false
	} else {
		return true
	}
}

//PregReplace 正则替换
//arr 正则表达式数组
//repl 替换的内容，可以是字符串，也可以是字符串数组
//src 要替换的字符串
func PregReplace(arr []string, repl interface{}, src string) string {
	for i, s := range arr {
		if value, ok := repl.(string); ok {
			src = regexp.MustCompile(s).ReplaceAllString(src, value)
		} else if value, ok := repl.([]string); ok {
			src = regexp.MustCompile(s).ReplaceAllString(src, value[i])
		}
	}
	return src
}

func HideString(str string, starNum int,starStr ...string) string {
	hLen := len([]rune(str))
	min := int(math.Floor(float64(hLen) / 3))
	if starNum > 0 {
		min = starNum
	}
	star := hLen - (min * 2)
	re, _ := regexp.Compile(fmt.Sprintf("(.{%v}?)(.{%v}?)(.+?)", min, star))
	if len(starStr) > 0 {
		newStr := re.ReplaceAllString(str, fmt.Sprintf("$1%v$3", starStr[0]))
		return newStr
	}
	newStr := re.ReplaceAllString(str, "$1****$3")
	return newStr
}

//ChunkSplit 字符串间隔插入字符
func ChunkSplit(body string, chunklen uint, end string) string {
	if end == "" {
		end = "\r\n"
	}
	runes, erunes := []rune(body), []rune(end)
	l := uint(len(runes))
	if l <= 1 || l < chunklen {
		return body + end
	}
	ns := make([]rune, 0, len(runes)+len(erunes))
	var i uint
	for i = 0; i < l; i += chunklen {
		if i+chunklen > l {
			ns = append(ns, runes[i:]...)
		} else {
			ns = append(ns, runes[i:i+chunklen]...)
		}
		ns = append(ns, erunes...)
	}
	return string(ns)
}

//FilterNumber 过滤掉字符串中的数字
func FilterNumber(str string) string {
	reg := regexp.MustCompile(`\d+`)
	return reg.ReplaceAllString(str, "")
}
//FilterNoNumber 过滤掉字符串中的非数字
func FilterNoNumber(str string) string {
	reg := regexp.MustCompile(`\D+`)
	return reg.ReplaceAllString(str, "")
}

func Price(price float64, decimals int) string {
	var priceStr = fmt.Sprintf("%."+fmt.Sprintf("%d", decimals)+"f", price)
 if strings.Contains(priceStr, "."+strings.Repeat("0", decimals)) {
	 return fmt.Sprintf("%.0f", price)
 }else {
	 return priceStr
 }
}
//中文字数
func StrLen(str string) int {
	return len([]rune(str))
}