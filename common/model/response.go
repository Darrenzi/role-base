package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	c    *gin.Context
	Code int
	Desc string
	Data interface{}
}

func (r *Response) Success(data interface{}) {
	r.Code = SUCCESS
	r.Desc = GetDesc(SUCCESS)
	r.Data = data

	r.c.JSON(http.StatusOK, r)
}

func (r *Response) Error() {
	r.Code = ERROR
	r.Desc = GetDesc(ERROR)

	r.c.JSON(http.StatusInternalServerError, r)
}

func (r *Response) Invalid() {
	r.Code = INVALID_PARAMS
	r.Desc = GetDesc(INVALID_PARAMS)

	r.c.JSON(http.StatusBadRequest, r)
}

func (r *Response) InvalidWithDesc(desc string) {
	r.Code = INVALID_PARAMS
	r.Desc = desc

	r.c.JSON(http.StatusBadRequest, r)
}

func (r *Response) ErrorWithDesc(desc string) {
	r.Code = ERROR
	r.Desc = desc

	r.c.JSON(http.StatusInternalServerError, r)
}

func (r *Response) UnauthorizedAuthFail() {
	r.Code = UnauthorizedAuthFail
	r.Desc = GetDesc(UnauthorizedAuthFail)

	r.c.JSON(http.StatusUnauthorized, r)
}

func JSON(httpCode int, errorCode int, c *gin.Context) {
	c.JSON(httpCode, Response{
		Code: errorCode,
		Desc: GetDesc(errorCode),
		Data: nil,
	})
}

func NewResponse(c *gin.Context) *Response {
	return &Response{c: c}
}
