package basepayload

import (
	"stepp-backend/src/constant"
)


type BasePayload struct {
	Code 		string 		`json:"code"`
	Description string 		`json:"description"`
	Data		interface{} `json:"data"`
}

func GetSuccessResponse(data interface{}) *BasePayload {
	return &BasePayload{
		Code: constant.SUCCESS_CODE,
		Description: constant.SUCCESS_DESCRIPTION,
		Data: data,
	}
}