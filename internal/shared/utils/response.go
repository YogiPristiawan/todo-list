package utils

import (
	"log"
	"net/http"

	"github.com/pkg/errors"
)

// CommonResult provides data struct
// that identifies how response should be
// returned, either success or fail
type CommonResult struct {
	code    int    `json:"-"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// SetResponse set the response code, and error if exists
func (c *CommonResult) SetResponse(code int, err error) {
	c.code = code

	if code == 200 || code == 201 {
		c.Status = "Success"
	}

	if code > 201 && code < 400 {
		c.Status = http.StatusText(code)
	}

	if code >= 400 && code < 500 { // client error
		c.Status = http.StatusText(code)
		c.Message = err.Error()
		return
	}

	if code >= 500 { // server error
		// send to logger
		go func() {
			log.Println(errors.WithStack(err))
		}()

		c.Status = http.StatusText(code)
		c.Message = "something wrong"
	}
}

// GetCode return response status code
func (c *CommonResult) GetCode() int {
	return c.code
}

type BaseResponse struct {
	CommonResult
	Data interface{} `json:"data"`
}

type BaseResponseArray[T interface{}] struct {
	CommonResult
	Data []T `json:"data"`
}
