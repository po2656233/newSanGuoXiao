package hints

import (
	"github.com/po2656233/superplace/components/gin"
	sgxCode "github.com/po2656233/superplace/const/code"
)

type Result struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewDataResult(code int32) *Result {
	result := &Result{
		Code:    code,
		Message: sgxCode.GetMessage(code),
		Data:    []string{},
	}
	return result
}

func (p *Result) SetCode(code int32) {
	p.Code = code
	p.Message = sgxCode.GetMessage(code)
}

func RenderResult(c *superGin.Context, statusCode int32, data ...interface{}) {
	result := NewDataResult(statusCode)
	if len(data) > 0 {
		result.Data = data[0]
	}
	c.JSON200(result)
}
