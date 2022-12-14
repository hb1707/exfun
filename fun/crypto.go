package fun

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"log"
	"math/rand"
	"net/url"
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

//UrlEncode W3C
func UrlEncode(uDec string) string {
	uEnc := url.QueryEscape(uDec)
	return uEnc
}

//UrlDecode W3C
func UrlDecode(uEnc string) string {
	uDec, err := url.QueryUnescape(uEnc)
	if err != nil {
		return ""
	} else {
		return uDec
	}
}

//UrlEncodePlus //RFC 2396
func UrlEncodePlus(uDec string) string { //RFC 2396
	uEnc := url.QueryEscape(uDec)
	return strings.Replace(uEnc, "+", "%20", -1)
}

//UrlDecodePlus //RFC 2396
func UrlDecodePlus(uEnc string) string { //RFC 2396
	uDec, err := url.QueryUnescape(uEnc)
	if err != nil {
		return ""
	} else {
		return strings.Replace(uDec, "%20", "+", -1)
	}
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

//Base64urlencode
func Base64UrlEncode(uDec string) string {
	uEnc := base64.URLEncoding.EncodeToString([]byte(uDec))
	return uEnc
}

//Base64urldecode
func Base64UrlDecode(uEnc string) string {
	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		return ""
	} else {
		return string(uDec)
	}
}

//urlencode处理加号
func Base64UrlEncodePlus(str string) string {
	return strings.Replace(Base64UrlEncode(str), "+", "%20", -1)
}

//urldecode处理加号
func Base64UrlDecodePlus(str string) string {
	return Base64UrlDecode(strings.Replace(str, "%20", "+", -1))
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
//生成随机字母字符串
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

//可逆加密
func Encrypt(text string, key string) []byte {
	var result string
	text = StrReverse(Base64Encode(text, true))
	for i := 0; i < len(text); i++ {
		result += string(text[i] ^ key[i%len(key)])
	}
	return []byte(result)
}

//可逆解密
func Decrypt(textByte []byte, key string) string {
	text := string(textByte)
	var result string
	for i := 0; i < len(text); i++ {
		result += string(text[i] ^ key[i%len(key)])
	}
	result = Base64Decode(StrReverse(result), true)
	return result
}

//AesEncrypt AES加密
func AesEncrypt(orig, key string) string {
	origData := []byte(orig)
	k := []byte(key)

	block, err := aes.NewCipher(k)
	if err != nil {
		return ""
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted)
}

//AesDecrypt AES解密
func AesDecrypt(crypto, key string) string {
	cryptoByte, _ := base64.StdEncoding.DecodeString(crypto)
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return ""
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	origData := make([]byte, len(crypto))
	blockMode.CryptBlocks(origData, cryptoByte)
	origData = PKCS5UnPadding(origData)
	return string(origData)
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
