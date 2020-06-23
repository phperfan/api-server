package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	tnet "github.com/toolkits/net"
)

//GenUUID 生成一个UUID
func GenUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

// GetReqID 获取请求中的request_id
func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-ID")
	if !ok {
		return ""
	}
	if requestID, ok := v.(string); ok {
		return requestID
	}
	return ""
}

var (
	once     sync.Once
	clientIP = "127.0.0.1"
)

// GetLocalIP 获取本地内网IP
func GetLocalIP() string {
	once.Do(func() {
		ips, _ := tnet.IntranetIP()
		if len(ips) > 0 {
			clientIP = ips[0]
		} else {
			clientIP = "127.0.0.1"
		}
	})
	return clientIP
}

// Md5 字符串转md5
func Md5(str string) (string, error) {
	h := md5.New()

	_, err := io.WriteString(h, str)
	if err != nil {
		return "", err
	}

	// 注意：这里不能使用string将[]byte转为字符串，否则会显示乱码
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
