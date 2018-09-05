package datamodels

type Rating struct {
	ID       int64  `json:"userId" form:"id"`
	Username string `json:"username" form:"username"`
}
//申请的状态：
//0 :审核未通过
//1：信誉背书通过
//-1: 信誉背书制未通过
//2：实名制通过
//-2:实名制未通过
//3：审核通过
type RatingStatus struct {
	Status int `json:"status"`
}

//检查信誉码是否有效
type CreditCode struct {
	CreditCode string `json:"creditCode"`
}

type ApplyRater struct {
	step int `json:"step"`
	codes []string `json:"codes"` //
	certificates []string `json:"certificates"` //证件照数组
}

type PageUser struct {
	UserId int64 `json:"userId"`
	StartIndex int `json:"startIndex"` //获取数据的起始位置
	RequestSize int `json:"requestSize"` //获取数据的个数
}

type Project struct {
	ProjectId int64 `json:"projectId"`
	Name string `json:"name"`
	Logo string `json:"logo"`
	Token string `json:"token"`
	Tags []string `json:"tags"`
}
type PageCategory struct {
	Category string `json:"category"`
	StartIndex int `json:"startIndex"` //获取数据的起始位置
	RequestSize int `json:"requestSize"` //获取数据的个数
}

type List struct {
	List interface{} `json:"list"`
}

type FocusRaterId struct {
	RaterId int64 `json:"raterId"`
	Focus int `json:"focus"` //0:取消关注 	1：关注
}

//检查用户是否可以评级
type CanRate struct {
	UserId int64 `json:"userId"`
	ProjectId int64 `json:"projectId"`
	CanRate int `json:"canRate"`//1:可以 	0：不可以
}