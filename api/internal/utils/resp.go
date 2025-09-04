package utils

import (
	"blog/api/internal/types"
)

func NewSuccessResp() *types.BaseResp {
	return &types.BaseResp{
		Code:    0,
		Message: "success",
	}
}
func NewRespWithMessage(code int, message string) *types.BaseResp {
	return &types.BaseResp{
		Code:    code,
		Message: message,
	}
}

func NewErrRespWithCode(code int) *types.BaseResp {
	return &types.BaseResp{
		Code:    code,
		Message: ErrorCodeMessages[code],
	}
}

func NewErrRespWithMessage(code int, message string) *types.BaseResp {
	return &types.BaseResp{
		Code:    code,
		Message: ErrorCodeMessages[code] + ": " + message,
	}
}
