package middleware

import (
	"api-server/pkg/log"
	"api-server/pkg/util"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

//type bodyLogWriter struct {
//	gin.ResponseWriter
//	body *bytes.Buffer
//}
//
//func (w bodyLogWriter) Write(b []byte) (int, error) {
//	w.body.Write(b)
//	return w.ResponseWriter.Write(b)
//}

// Logging is a middleware function that logs the each request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		// Continue.
		c.Next()

		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)

		log.Hl.WithFields(logrus.Fields{
			"request": c.Request,
			"path":    path,
			"latency": float64(latency) / 1000000,
			"ip":      util.GetLocalIP(),
		}).Info("")

	}
}
