package api

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
	"github.com/happyfromtbq/ratingchain-web/services"
	"github.com/kataras/iris/sessions"
	"github.com/happyfromtbq/ratingchain-web/datamodels"
)

type ProjectController struct {
	// context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding)
	// and the Session which depends on the current context (dynamic binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	Service services.RatingService

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session
}

//查看评级者对项目评级维度
func (c *ProjectController) PostGetratervsproject() mvc.Result{
	var m datamodels.CanRate
	if err := c.Ctx.ReadJSON(&m);
		err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return mvc.Response{
			Err: err,
		}
	}
	var v1 = datamodels.ProjectTag{
		Name:"架构",
		Value:6.5,
		Level:1,
	}
	var v2 = datamodels.ProjectTag{
		Name:"安全性",
		Value:6.5,
		Level:1,
	}
	var apiMsg = SuccessApiMsg
	apiMsg.ResponseData = datamodels.List{[]datamodels.ProjectTag{v1,v2}}
	return mvc.Response{
		Object: apiMsg,
	}
}

//获取用户关注的项目列表
func (c *ProjectController) PostGetfocusprojects()  mvc.Result{
	var m datamodels.PageUser
	if err := c.Ctx.ReadJSON(&m);
		err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return mvc.Response{
			Err: err,
		}
	}

	var p1 = datamodels.Project{
		ProjectId:343432,
		Name:"初链",
		Logo:"http://www.truechain.pro/logo.png",
		Token:"TRUE",
		Tags:[]string{"公链", "基础链"},
		Raters:100,
		Score:7.5,
		Rater:90,
	}
	var p2 = datamodels.Project{
		ProjectId:4545,
		Name:"以太坊",
		Logo:"http://www.eth.io/logo.png",
		Token:"ETH",
		Tags:[]string{"公链", "基础链"},
		Raters:500,
		Score:9,
		Rater:490,
	}
	var apiMsg = SuccessApiMsg
	apiMsg.ResponseData = datamodels.List{[]datamodels.Project{p1,p2}}
	return mvc.Response{
		Object: apiMsg,
	}
}

//获取项目分类
func (c *ProjectController) PostGetprojectcategory() mvc.Result{
	var apiMsg = SuccessApiMsg
	apiMsg.ResponseData = datamodels.List{[]string{"架构","安全","内容"}}
	return mvc.Response{
		Object: apiMsg,
	}
}

//获取项目分类
func (c *ProjectController) PostGetcategoryprojects() mvc.Result{
	var m datamodels.PageCategory
	if err := c.Ctx.ReadJSON(&m);
		err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return mvc.Response{
			Err: err,
		}
	}

	var p1 = datamodels.Project{
		ProjectId:343432,
		Name:"初链",
		Logo:"http://www.truechain.pro/logo.png",
		Token:"TRUE",
		Tags:[]string{"公链", "基础链"},
		Raters:100,
		Score:7.5,
		Rater:90,
	}
	var p2 = datamodels.Project{
		ProjectId:4545,
		Name:"以太坊",
		Logo:"http://www.eth.io/logo.png",
		Token:"ETH",
		Tags:[]string{"公链", "基础链"},
		Raters:500,
		Score:9,
		Rater:490,
	}
	var apiMsg = SuccessApiMsg
	apiMsg.ResponseData = datamodels.List{[]datamodels.Project{p1,p2}}
	return mvc.Response{
		Object: apiMsg,
	}
}

//获取项目的基本信息
func (c *ProjectController) PostGetprojectdetail() mvc.Result{
	var m datamodels.CanRate
	if err := c.Ctx.ReadJSON(&m);
		err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return mvc.Response{
			Err: err,
		}
	}
	//var fr = datamodels.FocusRaterId{
	//	RaterId:111,//userId
	//	Focus:1,
	//}
	var p1 = datamodels.Project{
		ProjectId:343432,
		Name:"初链",
		Logo:"http://www.truechain.pro/logo.png",
		Token:"TRUE",
		Tags:[]string{"公链", "基础链"},
		Description:"详细描述很多内容",
		Raters:100,
		Score:7.5,
		Rater:90,
		Focus:1,
	}
	var apiMsg = SuccessApiMsg
	apiMsg.ResponseData = p1
	return mvc.Response{
		Object: apiMsg,
	}
}

//获取项目的基本信息
func (c *ProjectController) PostGetstatisticsscore() mvc.Result {
	var m datamodels.CanRate
	if err := c.Ctx.ReadJSON(&m);
		err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return mvc.Response{
			Err: err,
		}
	}
	var p datamodels.ProjectTagScore
	p.TotalScore = 78.5
	var pt1 = datamodels.ProjectTag{
		Name:"架构",
		Value:76.5,
	}
	var pt2 = datamodels.ProjectTag{
		Name:"安全性",
		Value:85,
	}
	p.Dimensions = []datamodels.ProjectTag{pt1,pt2}
	var apiMsg = SuccessApiMsg
	apiMsg.ResponseData = p
	return mvc.Response{
		Object: apiMsg,
	}
}

//获取参与项目的评级者
func (c *ProjectController) PostGetprojectraters() mvc.Result {
	var m datamodels.ProjectPage
	if err := c.Ctx.ReadJSON(&m);
		err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return mvc.Response{
			Err: err,
		}
	}

	var apiMsg = SuccessApiMsg
	var userInfo1 = datamodels.UserInfo{
		UserId:11,
		Username:"ASA",
		Level:3,
		Credit:4.5,
		Tags:[]string{"架构","安全","稳定"},
		Project:12,
	}
	var userInfo2= datamodels.UserInfo{
		UserId:12,
		Username:"bbb",
		Level:3,
		Credit:4.5,
		Tags:[]string{"架构","安全","稳定"},
		Project:12,
	}

	apiMsg.ResponseData = datamodels.List{[]datamodels.UserInfo{userInfo1,userInfo2}}
	return mvc.Response{
		Object: apiMsg,
	}
}

//评级时获取评级指标
func (c *ProjectController) PostGetratedimensions() mvc.Result {
	var v1 = datamodels.ProjectTag{
		Name:"架构",
		Value:6.5,
		Level:1,
	}
	var v2 = datamodels.ProjectTag{
		Name:"安全性",
		Value:6.5,
		Level:1,
	}
	var apiMsg = SuccessApiMsg
	apiMsg.ResponseData = datamodels.List{[]datamodels.ProjectTag{v1,v2}}
	return mvc.Response{
		Object: apiMsg,
	}
}

//提交项目评级
func (c *ProjectController) PostSubmitratingscores() mvc.Result {
	var m datamodels.ProjectTagScore
	if err := c.Ctx.ReadJSON(&m);
		err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return mvc.Response{
			Err: err,
		}
	}
	return mvc.Response{
		Object: SuccessApiMsg,
	}
}

//关注项目
func (c *ProjectController) PostFocusproject() mvc.Result {
	var m datamodels.ProjectFocus
	if err := c.Ctx.ReadJSON(&m);
		err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return mvc.Response{
			Err: err,
		}
	}
	//todo 关注或取消关注项目
	return mvc.Response{
		Object: SuccessApiMsg,
	}
}

