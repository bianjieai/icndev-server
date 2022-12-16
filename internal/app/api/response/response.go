package response

import (
	"github.com/bianjieai/icndev-server/internal/app/errors"
	"github.com/bianjieai/icndev-server/internal/app/model/vo"
)

func Build(data interface{}) vo.BaseResponse {
	return vo.BaseResponse{
		Code: 0,
		Data: data,
	}
}

func BuildWithMessage(message string, data interface{}) vo.BaseResponse {
	return vo.BaseResponse{
		Code:    0,
		Message: message,
		Data:    data,
	}
}

func BuildWithPage(data interface{}, pageVO vo.PageVO, total int) vo.BaseResponseWithPage {
	return vo.BaseResponseWithPage{
		Code:  0,
		Data:  data,
		Page:  pageVO.Page,
		Size:  pageVO.Size,
		Total: total,
	}
}

func BuildErr(err errors.Error) vo.BaseResponse {
	return vo.BaseResponse{
		Code:    err.Code(),
		Message: err.Msg(),
	}
}

func BuildErrCode(errCode int) vo.BaseResponse {
	return vo.BaseResponse{
		Code:    errCode,
		Message: errors.GetMsg(errCode),
	}
}

func BuildSubscribed(err errors.Error) vo.BaseResponse {
	return vo.BaseResponse{
		Code:    err.Code(),
		Message: "Subscribed",
	}
}
