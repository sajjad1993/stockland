package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (writer bodyLogWriter) Write(body []byte) (int, error) {
	writer.body.Write(body)
	return writer.ResponseWriter.Write(body)
}

func (writer bodyLogWriter) WriteString(body string) (int, error) {
	writer.body.WriteString(body)
	return writer.ResponseWriter.WriteString(body)
}
