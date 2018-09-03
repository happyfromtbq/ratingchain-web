package api

import (
	"github.com/happyfromtbq/ratingchain-web/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/mvc"
	"github.com/happyfromtbq/ratingchain-web/datamodels"
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

// PostRegister handles POST:
func (c *UserController) PostGetRaterStatus() mvc.Result {
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

