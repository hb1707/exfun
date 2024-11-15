package fun

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"log"
	mathRand "math/rand"
	"net/url"
	"strings"
	"time"
	"unsafe"
)

// SHA256 sha256
func SHA256(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// MD5 md5
func MD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// UrlEncode W3C
func UrlEncode(uDec string) string {
	uEnc := url.QueryEscape(uDec)
	return uEnc
}

// UrlDecode W3C
func UrlDecode(uEnc string) string {
	uDec, err := url.QueryUnescape(uEnc)
	if err != nil {
		return ""
	} else {
		return uDec
	}
}

// UrlEncodePlus //RFC 2396
func UrlEncodePlus(uDec string) string { //RFC 2396
	uEnc := url.QueryEscape(uDec)
	return strings.Replace(uEnc, "+", "%20", -1)
}

// UrlDecodePlus //RFC 2396
func UrlDecodePlus(uEnc string) string { //RFC 2396
	uDec, err := url.QueryUnescape(uEnc)
	if err != nil {
		return ""
	} else {
		return strings.Replace(uDec, "%20", "+", -1)
	}
}

// Base64Encode base64_encode
func Base64Encode(sDec string, security bool) string {
	var sEnc string
	if security {
		sEnc = base64.RawURLEncoding.EncodeToString([]byte(sDec))
	} else {
		sEnc = base64.StdEncoding.EncodeToString([]byte(sDec))
	}
	return sEnc
}

// Base64Decode base64_decode
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

// Base64UrlEncode Base64urlencode
func Base64UrlEncode(uDec string) string {
	uEnc := base64.URLEncoding.EncodeToString([]byte(uDec))
	return uEnc
}

// Base64UrlDecode Base64urldecode
func Base64UrlDecode(uEnc string) string {
	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		return ""
	} else {
		return string(uDec)
	}
}

// Base64UrlEncodePlus urlencode处理加号
func Base64UrlEncodePlus(str string) string {
	return strings.Replace(Base64UrlEncode(str), "+", "%20", -1)
}

// Base64UrlDecodePlus urldecode处理加号
func Base64UrlDecodePlus(str string) string {
	return Base64UrlDecode(strings.Replace(str, "%20", "+", -1))
}

// StrReverse 字符串逆序
func StrReverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

var src = mathRand.NewSource(time.Now().UnixNano())

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits
)

// RandString https://www.flysnow.org/2019/09/30/how-to-generate-a-random-string-of-a-fixed-length-in-go.html
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

// Encrypt 可逆加密
func Encrypt(text string, key string) []byte {
	var result string
	text = StrReverse(Base64Encode(text, true))
	for i := 0; i < len(text); i++ {
		result += string(text[i] ^ key[i%len(key)])
	}
	return []byte(result)
}

// Decrypt 可逆解密
func Decrypt(textByte []byte, key string) string {
	text := string(textByte)
	var result string
	for i := 0; i < len(text); i++ {
		result += string(text[i] ^ key[i%len(key)])
	}
	result = Base64Decode(StrReverse(result), true)
	return result
}

// AesEncrypt AES加密
func AesEncrypt(orig, keyStr string) string {
	//key, err := base64.StdEncoding.DecodeString(keyStr)
	key := []byte(keyStr)
	origData := []byte(orig)
	//k := []byte(key)
	//iv := make([]byte, aes.BlockSize)
	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}
	blockSize := block.BlockSize()
	if len(key) < blockSize {
		return ""
	}
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted)
}

// AesDecrypt AES解密
func AesDecrypt(crypto, keyStr string) string {
	//key, err := base64.StdEncoding.DecodeString(keyStr)
	key := keyStr
	cryptoByte, _ := base64.StdEncoding.DecodeString(crypto)
	//iv := make([]byte, aes.BlockSize)
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return ""
	}
	blockSize := block.BlockSize()
	if len(key) < blockSize {
		return ""
	}
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	origData := make([]byte, len(cryptoByte))
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

// Rsa2Sign RSA2签名
func Rsa2Sign(data string, privateKey []byte) string {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return ""
	}
	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return ""
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(signature)
}
