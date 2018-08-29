package api

type ApiMsg struct {
	Code string  `json:"code" form:"code"`
	Message string `json:"message" form:"message"`
	ResponseData interface{} `json:"responseData,omitempty" form:"responseData"`
}

var SuccessApiMsg =  ApiMsg{
	Code: "0",
	Message:"操作成功",
}

var FailApiMsg = ApiMsg{
	Code:"-1",
	Message:"操作失败",
}

