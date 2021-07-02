package fun

import (
	"net/http"
)

func ClientIP(r *http.Request, isCDN bool) string {
	//ip = SERVER.REMOTE_ADDR
	var (
		matches    []string
		matchesAll [][]string
	)
	httpVia := r.Header["HTTP_VIA"][0]                  //代理/节点 IP
	remoteAddr := r.RemoteAddr                          //最后一个代理服务器 IP，最真实，但可能被隐藏，也可能只是节点IP
	httpClientIP := r.Header["HTTP_CLIENT_IP"][0]       //非标准，不是所有的负载均衡节点都带
	httpXForwardedFor := r.Header["X-Forwarded-For"][0] //多个代理服务器 IP
	ip := httpVia

	if ip == "" {
		ip = remoteAddr
	}
	if ip != "" && PregMatch(`/^([0-9]{1,3}\.){3}[0-9]{1,3}/`, ip, nil) && isCDN {
		return ip
	}
	if httpClientIP != "" && PregMatch(`/^([0-9]{1,3}\.){3}[0-9]{1,3}/`, httpClientIP, &matches) {
		ip = httpClientIP
	} else if httpXForwardedFor != "" && PregMatchAll(`#\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}#s`, httpXForwardedFor, &matchesAll) {

		for _, xip := range matchesAll[0] {
			if !PregMatch(`#^(10|172\.16|192\.168)\.#`, xip, nil) {
				ip = xip
				break
			}
		}
	}
	return ip
}
