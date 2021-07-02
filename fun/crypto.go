package fun

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"log"
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

//sha256
func SHA256(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//md5
func MD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//base64_encode
func Base64Encode(sDec string, security bool) string {
	var sEnc string
	if security {
		sEnc = base64.RawURLEncoding.EncodeToString([]byte(sDec))
	} else {
		sEnc = base64.StdEncoding.EncodeToString([]byte(sDec))
	}
	return sEnc
}

//base64_decode
func Base64Decode(sEnc string, security bool) string {
	var sDecByte []byte
	var err error
	if security {
		sDecByte, err = base64.RawURLEncoding.DecodeString(sEnc)
	} else {
		sDecByte, err = base64.StdEncoding.DecodeString(sEnc)
	}
	if err != nil {
		log.Println(err.Error())
		return ""
	} else {
		return string(sDecByte)
	}

}

//urlencode
func UrlEncode(uDec string) string {
	uEnc := base64.URLEncoding.EncodeToString([]byte(uDec))
	return uEnc
}

//urldecode
func UrlDecode(uEnc string) string {
	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		return ""
	} else {
		return string(uDec)
	}
}

//rawurlencode
func RawUrlEncode(str string) string {
	return strings.Replace(UrlEncode(str), "+", "%20", -1)
}

//rawurldecode
func RawUrlDecode(str string) string {
	return UrlDecode(strings.Replace(str, "%20", "+", -1))
}

//字符串逆序
func StrReverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits
)

//https://www.flysnow.org/2019/09/30/how-to-generate-a-random-string-of-a-fixed-length-in-go.html
func RandString(n int) string {
	var b = make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
