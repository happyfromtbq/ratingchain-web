// file: controllers/user_controller.go

package api

import (
	"github.com/happyfromtbq/ratingchain-web/datamodels"
	"github.com/happyfromtbq/ratingchain-web/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

// UserController is our /user controller.
// UserController is responsible to handle the following requests:
// GET  			/user/register
// POST 			/user/register
// GET 				/user/login
// POST 			/user/login
// GET 				/user/me
// All HTTP Methods /user/logout
type UserController struct {
	// context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding)
	// and the Session which depends on the current context (dynamic binding).
	Ctx iris.Context

	// Our UserService, it's an interface which
	// is binded from the main application.
	Service services.UserService

	// Session, binded using dependency injection from the main.go.
	Session *sessions.Session
}

const userIDKey = "UserID"

func (c *UserController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(userIDKey, 0)
	return userID
}

func (c *UserController) isLoggedIn() bool {
	return c.getCurrentUserID() > 0
}

func (c *UserController) logout() {
	c.Session.Destroy()
}



// PostRegister handles POST: http://localhost:8080/api/user/register.
func (c *UserController) PostRegister() mvc.Result {
	// get firstname, username and password from the form.
	//var (
	//	firstname = c.Ctx.FormValue("firstname")
	//	username  = c.Ctx.FormValue("username")
	//	password  = c.Ctx.FormValue("password")
	//)
	//var h_token = c.Ctx.GetHeader("token")
	var c_device = c.Ctx.GetHeader("device")
	if(c_device == ""){
		var msg ApiMsg
		msg.Code = "300"
		msg.Message = "缺少device信息"
		msg.ResponseData = nil
		return mvc.Response{
			Object: msg,
		}
	}
	var u datamodels.User

	if err := c.Ctx.ReadJSON(&u); err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return mvc.Response{
			Err: err,
		}
	}

	// create the new user, the password will be hashed by the service.
	u, err := c.Service.Create(u.Password, u)

	var reUser datamodels.User
	reUser = datamodels.User{
		Username:  u.Username,
		Token: u.Token,
		ID:u.ID,
	}

	// set the user's id to this session even if err != nil,
	// the zero id doesn't matters because .getCurrentUserID() checks for that.
	// If err != nil then it will be shown, see below on mvc.Response.Err: err.
	c.Session.Set(userIDKey, u.ID)

	//build response Api Msg
	var msg ApiMsg
	msg.Code = "200"
	msg.Message = "创建账号成功"
	msg.ResponseData = reUser

	return mvc.Response{
		// if not nil then this error will be shown instead.
		Err: err,
		// redirect to /user/me.
		//Path: "/user/me",

		//response Json
		Object: msg,
		// When redirecting from POST to GET request you -should- use this HTTP status code,
		// however there're some (complicated) alternatives if you
		// search online or even the HTTP RFC.
		// Status "See Other" RFC 7231, however iris can automatically fix that
		// but it's good to know you can set a custom code;
		// Code: 303,
	}

}

// PostLogin handles POST: http://localhost:8080/user/login.
func (c *UserController) PostLogin() mvc.Result {
	var u datamodels.User

	if err := c.Ctx.ReadJSON(&u);
	err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return mvc.Response{
			Err: err,
		}
	}

	u, found := c.Service.GetByUsernameAndPassword(u.Username, u.Password)

	if !found {
		var msg ApiMsg
		msg.Code = "200"
		msg.Message = "账号不存在或密码错误"
		msg.ResponseData = nil
		return mvc.Response{
			Object: msg,
		}
	}

	c.Session.Set(userIDKey, u.ID)

	var reUser datamodels.User
	reUser = datamodels.User{
		Username:  u.Username,
		Token: u.Token,
		ID:u.ID,
	}
	var msg ApiMsg
	msg.Code = "200"
	msg.Message = "用户登录成功"
	msg.ResponseData = reUser

	return mvc.Response{
		Object: msg,
	}
}

// PostLogin handles POST: /logout.
func (c *UserController) PostLogout() mvc.Result {
	var token string = c.Ctx.FormValue("token")
	if(token == ""){
		var msg ApiMsg
		msg.Code = "200"
		msg.Message = "token不存在"
		msg.ResponseData = nil
		return mvc.Response{
		Object: msg,
		}
	}
	if c.isLoggedIn() {
		c.logout()
	}
	var msg ApiMsg
	msg.Code = "200"
	msg.Message = "用户注销登录成功"

	return mvc.Response{
		Object: msg,
	}
}

func (c *UserController) PostChangePassword() mvc.Result {
	var u datamodels.User
	if err := c.Ctx.ReadJSON(&u);
	err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.WriteString(err.Error())
		return mvc.Response{
			Err: err,
		}
	}
	var msg ApiMsg
	msg.Code = "200"
	msg.Message = "用户修改密码成功"

	return mvc.Response{
		Object: msg,
	}
}

// GetMe handles GET: http://localhost:8080/user/me.
func (c *UserController) GetMe() mvc.Result {
	if !c.isLoggedIn() {
		// if it's not logged in then redirect user to the login page.
		return mvc.Response{Path: "/user/login"}
	}

	u, found := c.Service.GetByID(c.getCurrentUserID())
	if !found {
		// if the  session exists but for some reason the user doesn't exist in the "database"
		// then logout and re-execute the function, it will redirect the client to the
		// /user/login page.
		c.logout()
		return c.GetMe()
	}

	return mvc.View{
		Name: "user/me.html",
		Data: iris.Map{
			"Title": "Profile of " + u.Username,
			"User":  u,
		},
	}
}

