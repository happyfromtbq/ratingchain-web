package api

type ApiMsg struct {
	Code string  `json:"code" form:"code"`
	Message string `json:"message" form:"message"`
	ResponseData interface{} `json:"responseData,omitempty" form:"responseData"`
}
