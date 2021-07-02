package curl

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const AccessToken = "Access-Token"
const ApplicationJson = "application/json"
const ApplicationXml = "application/xml;charset=utf-8"
const TlsTimeout = time.Duration(20 * time.Second)
const HttpTimeout = time.Duration(35 * time.Second)

type Config struct {
	Headers map[string]string
}

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
	var data = url.Values{}
	for k, v := range param {
		data.Add(k, v)
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

func (c *Config) POSTJSON(httpUrl string, params []byte) ([]byte, http.Header) {

	var jsonStr = []byte(params)
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
	}
	client := &http.Client{Transport: transCfg}
	req, _ := http.NewRequest("POST", httpUrl, bytes.NewBuffer(jsonStr))
	//  组装http请求头
	req.Header.Set("Content-Type", "application/json")
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil, nil
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	respHeader := resp.Header
	return respBody, respHeader
}
func (c *Config) DELETE(httpUrl string, param map[string]string) ([]byte, int, error) {
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
