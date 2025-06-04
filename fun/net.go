package fun

import (
	"net/http"
	"strings"
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

// Domain 从 URL 获取顶级域名
func Domain(url string, suffix string) string {
	// 移除协议部分
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")

	// 移除路径和查询参数
	if i := strings.Index(url, "/"); i > 0 {
		url = url[:i]
	}

	// 移除端口号
	if i := strings.Index(url, ":"); i > 0 {
		url = url[:i]
	}

	// 分割域名部分
	parts := strings.Split(url, ".")
	partsLen := len(parts)

	// 处理域名情况
	if partsLen <= 2 {
		// 例如: example.com
		return url
	} else {
		// 处理子域名情况，例如: sub.example.com
		// 对于某些特殊的二级域名，如 co.uk, com.cn 等，这里需要额外处理
		// 这里给出一个简化实现
		if strings.HasSuffix(url, suffix) {
			// 对于特殊二级域名，返回最后三段
			return parts[partsLen-3] + "." + parts[partsLen-2] + "." + parts[partsLen-1]
		}
		// 对于普通域名，返回最后两段
		return parts[partsLen-2] + "." + parts[partsLen-1]
	}
}
