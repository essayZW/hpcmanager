package response

import "github.com/gin-gonic/gin"

// Response 接口响应的数据
type Response struct {
	// code 接口返回值
	Code int `json:"code"`
	// data 接口返回数据
	Data interface{} `json:"data"`
	// status 操作的成功与否
	Status bool `json:"status"`
	// message 接口返回的消息
	Message string `json:"message"`
}

// Send 发送请求
func (resp *Response) Send(ctx *gin.Context) {
	ctx.JSON(200, resp)
}

// New 创建一个新的响应体
func New(code int, data interface{}, status bool, message string) *Response {
	return &Response{code, data, status, message}
}
