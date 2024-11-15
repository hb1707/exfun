package curl

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const AccessToken = "Access-Token"
const ApplicationJson = "application/json"
const ApplicationFile = "multipart/form-data"
const ApplicationXml = "application/xml;charset=utf-8"
const ApplicationForm = "application/x-www-form-urlencoded;charset=utf-8"
const ApplicationOctetStream = "application/octet-stream;charset=UTF-8"
const TlsTimeout = time.Duration(60 * time.Second)
const HttpTimeout = time.Duration(90 * time.Second)

type Config struct {
	Headers map[string]string
}

var IsDebug = make(map[string]bool, 0)
var DebugList = make(map[string][]string, 0)

var defaultConfig = &Config{
	Headers: map[string]string{"Content-Type": ApplicationJson},
}

var POST = defaultConfig.POST
var GET = defaultConfig.GET

func Web(httpUrl string, param map[string]string) string {
	var data = url.Values{}
	for k, v := range param {
		data.Add(k, v)
	}
	reqBody := data.Encode()
	return httpUrl + "?" + reqBody
}
func (c *Config) SetHeader(k, v string) {
	c.Headers[k] = v
}
func (c *Config) GET(httpUrl string, param map[string]string) ([]byte, int, error) {
	if httpUrl == "" {
		return nil, 0, fmt.Errorf("httpUrl is empty")
	}
	var data = url.Values{}
	for k, v := range param {
		data.Add(k, fmt.Sprintf("%v", v))
	}
	reqBody := data.Encode()
	client := &http.Client{}
	req, err := http.NewRequest("GET", httpUrl+"?"+reqBody, nil)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil, 0, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	return respBody, resp.StatusCode, nil
}
func (c Config) POST(httpUrl string, reqBody []byte) ([]byte, int, error) {
	if httpUrl == "" {
		return nil, 0, fmt.Errorf("httpUrl is empty")
	}
	transport := &http.Transport{
		TLSHandshakeTimeout: TlsTimeout,
		DisableKeepAlives:   true,
	}
	client := http.Client{
		Timeout:   HttpTimeout,
		Transport: transport,
	}
	req, err := http.NewRequest("POST", httpUrl, bytes.NewReader(reqBody))
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	c.Headers["Content-Length"] = fmt.Sprintf("%d", len(reqBody))
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	//core.Print("%+v",resp.Body)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return respBody, resp.StatusCode, nil
}
func (c Config) PUT(httpUrl string, reqBody []byte) ([]byte, int, error) {
	if httpUrl == "" {
		return nil, 0, fmt.Errorf("httpUrl is empty")
	}
	transport := &http.Transport{
		TLSHandshakeTimeout: TlsTimeout,
		DisableKeepAlives:   true,
	}
	client := http.Client{
		Timeout:   HttpTimeout,
		Transport: transport,
	}
	req, err := http.NewRequest("PUT", httpUrl, bytes.NewReader(reqBody))
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	//core.Print("%+v",resp.Body)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return respBody, resp.StatusCode, nil
}

func (c *Config) POSTFILE(httpUrl string, param map[string]string, file []byte) []byte {
	if httpUrl == "" {
		return nil
	}
	//  待合成文本
	var data = url.Values{}
	for k, v := range param {
		data.Add(k, v)
	}

	reqBody := data.Encode()

	//fmt.Printf("参数：%+v",reqBody)

	client := &http.Client{}

	req, _ := http.NewRequest("POST", httpUrl+"?"+reqBody, bytes.NewBuffer(file))
	//  组装http请求头
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	return respBody
}

func (c *Config) POSTJSON(httpUrl string, params []byte) ([]byte, http.Header, int) {
	if httpUrl == "" {
		return nil, nil, 0
	}
	var jsonStr = []byte(params)
	transCfg := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, // disable verify
		TLSHandshakeTimeout: TlsTimeout,
	}
	client := &http.Client{Transport: transCfg, Timeout: HttpTimeout}
	req, _ := http.NewRequest("POST", httpUrl, bytes.NewBuffer(jsonStr))
	//  组装http请求头
	req.Header.Set("Content-Type", ApplicationJson)
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil, nil, 0
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	respHeader := resp.Header
	c.AddDebug(req, httpUrl, params, respBody)
	return respBody, respHeader, resp.StatusCode
}
func (c *Config) DELETE(httpUrl string, param map[string]string) ([]byte, int, error) {
	if httpUrl == "" {
		return nil, 0, fmt.Errorf("httpUrl is empty")
	}
	var data = url.Values{}
	for k, v := range param {
		data.Add(k, v)
	}
	reqBody := data.Encode()
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", httpUrl+"?"+reqBody, nil)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil, 0, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	return respBody, resp.StatusCode, nil
}
func (c *Config) AddDebug(req *http.Request, httpUrl string, jsonStr []byte, respBody []byte) {
	if ok, exist := IsDebug[httpUrl]; exist && ok {
		header := ""
		for s, strings := range req.Header {
			header += fmt.Sprintf("%s:%s\n", s, strings)
		}
		if DebugList[httpUrl] == nil {
			DebugList[httpUrl] = make([]string, 0)
		}
		DebugList[httpUrl] = append([]string{fmt.Sprintf("POSTJSON: \n\nAPI:%s\n\nRequest:%s\n\nBody:%s", httpUrl, string(jsonStr), string(respBody))}, DebugList[httpUrl]...)
		if len(DebugList[httpUrl]) > 20 {
			DebugList[httpUrl] = DebugList[httpUrl][:20]
		}
	}
}

func (c *Config) POSTSEE(httpUrl string, params []byte, callback func([]byte)) int {
	if httpUrl == "" {
		return 0
	}
	var jsonStr = []byte(params)
	transCfg := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, // disable verify
		TLSHandshakeTimeout: TlsTimeout,
	}
	client := &http.Client{Transport: transCfg, Timeout: HttpTimeout}
	req, _ := http.NewRequest("POST", httpUrl, bytes.NewBuffer(jsonStr))
	//  组装http请求头
	req.Header.Set("Content-Type", ApplicationJson)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Connection", "keep-alive")
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return resp.StatusCode
	}
	defer resp.Body.Close()
	for {
		data := make([]byte, 1024)
		_, err := resp.Body.Read(data)
		if err != nil {
			return resp.StatusCode
		}
		if string(data) == "" || data == nil {
			break
		}
		callback(data)
	}
	return resp.StatusCode
}

func (c *Config) GETHtml(httpUrl string, param map[string]string) (io.ReadCloser, int, error) {
	if httpUrl == "" {
		return nil, 0, fmt.Errorf("httpUrl is empty")
	}
	var data = url.Values{}
	for k, v := range param {
		data.Add(k, fmt.Sprintf("%v", v))
	}
	reqBody := data.Encode()
	client := &http.Client{}
	req, err := http.NewRequest("GET", httpUrl+"?"+reqBody, nil)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil, 0, err
	}
	defer resp.Body.Close()
	return resp.Body, resp.StatusCode, nil
}
