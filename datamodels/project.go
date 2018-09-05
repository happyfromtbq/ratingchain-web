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
	TotalScore float64 `json:"totalScore"`
	Dimensions []ProjectTag `json:"dimensions"`
}
