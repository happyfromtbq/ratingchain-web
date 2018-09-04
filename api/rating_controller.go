package api

import (
	"github.com/happyfromtbq/ratingchain-web/services"
	"github.com/happyfromtbq/ratingchain-web/datamodels"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/mvc"

	"fmt"
)

type RatingController struct {
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


func (c *RatingController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(userIDKey, 0)
	return userID
}

func (c *RatingController) isLoggedIn() bool {
	return c.getCurrentUserID() > 0
}

func (c *RatingController) logout() {
	c.Session.Destroy()
}

// 获取评级者申请状态 handles POST:
func (c *RatingController) PostGetraterstatus() mvc.Result {
	if(c.isLoggedIn()){
		var apiMsg = SuccessApiMsg
		var ratingStatus = datamodels.RatingStatus{2}
		apiMsg.ResponseData = ratingStatus
		return mvc.Response{
			Object: apiMsg,
		}
	}
	return mvc.Response{
		Object: NoLoginApiMsg,
	}
}

//检查信誉码是否有效
func (c *RatingController) PostCheckcreditcode() mvc.Result {
	var i datamodels.CreditCode
	if err := c.Ctx.ReadJSON(&i);
		err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return mvc.Response{
			Err: err,
		}
	}
	var creditCode = i.CreditCode
	fmt.Print("参数是{}",creditCode)
	return mvc.Response{
		Object: SuccessApiMsg,
	}
}

//提交评级者申请信息
func (c *RatingController) PostApplyrater()  mvc.Result {
	var m datamodels.ApplyRater
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

//获取评级者参与的项目列表
func (c *RatingController) PostGetrateprojects()  mvc.Result {
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
		343432,
		"初链",
		"www.truechain.pro/logo.png",
		"TRUE",
		[]string{"公链", "基础链"},
	}
	var p2 = datamodels.Project{
		4545,
		"以太坊",
		"www.eth.io/logo.png",
		"ETH",
		[]string{"公链", "基础链"},
	}
	var apiMsg = SuccessApiMsg
	apiMsg.ResponseData = []datamodels.Project{p1,p2}
	return mvc.Response{
		Object: apiMsg,
	}
}

//获取评级者分类
func (c *RatingController) PostGetratercategory() mvc.Result{
	var apiMsg = SuccessApiMsg
	apiMsg.ResponseData = datamodels.List{[]string{"架构","安全","内容"}}
	return mvc.Response{
		Object: apiMsg,
	}
}

//根据分类获取评级者
func (c *RatingController) PostGetcategoryraters() mvc.Result{
	var m datamodels.PageCategory
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

//关注评级者
func (c *RatingController) PostFocusrater() mvc.Result{
	var m   datamodels.FocusRaterId
	if err := c.Ctx.ReadJSON(&m);
		err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return mvc.Response{
			Err: err,
		}
	}
	return mvc.Response{
		Object:SuccessApiMsg,
	}
}

//获取关注的评级者
func (c *RatingController) PostGetfocusraters() mvc.Result{
	var m datamodels.PageUser
	if err := c.Ctx.ReadJSON(&m);
		err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return mvc.Response{
			Err: err,
		}
	}
	//todo 待完成
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

