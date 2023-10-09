package response

import "github.com/gin-gonic/gin"

// CodeSuccess 成功代码
const CodeSuccess = 200

// CodeErr 错误代码
const CodeErr = 500

// MessageErr 错误消息
const MessageErr = "Internal Server Error"

// MessageSuccess 成功消息
const MessageSuccess = ""

// RespOK 响应结构体
type RespOK struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// RespData 响应结构体
type RespData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Ok 输出成功消息
func Ok(c *gin.Context, args ...string) {
	message := MessageSuccess
	if len(args) > 0 {
		message = args[0]
	}
	c.JSON(0, RespOK{
		Code:    CodeSuccess,
		Message: message,
	})
}

// Err 输出错误信息
func Err(c *gin.Context, args ...string) {
	message := MessageErr
	if len(args) > 0 {
		message = args[0]
	}
	c.JSON(0, RespOK{
		Code:    CodeErr,
		Message: message,
	})
	c.Abort()
}

// ErrData 输出错误信息
func ErrData(c *gin.Context, data interface{}, args ...string) {
	message := MessageErr
	if len(args) > 0 {
		message = args[0]
	}
	c.JSON(0, RespData{
		Code:    CodeErr,
		Message: message,
		Data:    data,
	})
	c.Abort()
}

// Data 通过指定的错误代码，输出错误信息
func Data(c *gin.Context, data interface{}) {
	c.JSON(0, RespData{
		Code:    CodeSuccess,
		Message: MessageSuccess,
		Data:    data,
	})
}
