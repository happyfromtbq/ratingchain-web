package datamodels

type Project struct {
	ProjectId int64 `json:"projectId"`
	Name string `json:"name"`
	Logo string `json:"logo"`
	Token string `json:"token"`
	Tags []string `json:"tags"`
	Description string `json:"description"`
	Raters int `json:"raters,omitempty"`
	Score float64 `json:"score,omitempty"`
	Rater int `json:"rater,omitempty"`
	//FocusRaterId FocusRaterId `json:"focusRaterId"` //1:被关注 0：未被关注
	Focus int `json:"focus"` //0:取消关注 	1：关注
}

//检查用户是否可以评级
type CanRate struct {
	UserId int64 `json:"userId,omitempty"`
	ProjectId int64 `json:"projectId,omitempty"`
	CanRate int `json:"canRate"`//1:可以 	0：不可以
}

type ProjectTag struct {
	Name string `json:"name"`
	Value float64 `json:"value"`
	Level int `json:"level"`
}

type ProjectTagScore struct {
	ProjectId int64 `json:"projectId"`
	TotalScore float64 `json:"totalScore"`
	Dimensions []ProjectTag `json:"dimensions"`
}

type ProjectPage struct {
	ProjectId int64 `json:"projectId"`
	StartIndex int `json:"startIndex"` //获取数据的起始位置
	RequestSize int `json:"requestSize"` //获取数据的个数
}

type ProjectFocus struct {
	ProjectId int64 `json:"projectId"`
	Focus int `json:"focus"`
}