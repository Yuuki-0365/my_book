package controller

import (
	"SmallRedBook/tool"
	"encoding/json"
)

func ErrorResponse(err error) tool.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return tool.Response{
			Status: 400,
			Msg:    "JSON类型不匹配",
			Error:  err.Error(),
		}
	}
	return tool.Response{
		Status: 400,
		Msg:    "参数错误",
		Error:  err.Error(),
	}
}
