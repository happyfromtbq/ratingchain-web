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