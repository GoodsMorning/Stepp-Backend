package basepayload

import (
	"stepp-backend/src/constant"
)

type BasePayload struct {
	Code        string      `json:"code"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func GetSuccessResponse(data interface{}) *BasePayload {
	return &BasePayload{
		Code:        constant.SUCCESS.CODE,
		Description: constant.SUCCESS.DESCRIPTION,
		Data:        data,
	}
}

func GetErrorResponse(errType constant.Response, data interface{}) *BasePayload {
	return &BasePayload{
		Code:        errType.CODE,
		Description: errType.DESCRIPTION,
		Data:        data,
	}
}
