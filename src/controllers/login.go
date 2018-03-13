package controllers

import (
	"github.com/go-steven/cms/src/controllers/base"
	"github.com/go-steven/cms/src/service"
	"github.com/go-steven/cms/src/util"
	"strconv"

	"github.com/astaxie/beego"
	"fmt"
)

type LoginController struct {
	base.BaseController
}

/**
进入登录页面
*/
func (c *LoginController) LoginView() {
	c.Show("common/loginPage.html")
}

/**
登陆
*/
func (c *LoginController) Login() {
	name := c.GetString("name")
	passwd := c.GetString("passwd")
	encodePwd := util.EncodeMessageMd5(passwd)
	beego.Debug(fmt.Sprintf("Login: name=%s, passwd=%s, encodePwd=%s", name, passwd, encodePwd))
	if user, err := service.UserService.Authentication(name, encodePwd); err != nil {
		c.JSON(err.Error())
	} else {
		token := strconv.FormatUint(user.Id, 10) + "|" + name + "|" + c.GetClientIp()
		beego.Debug(fmt.Sprintf("token: %s", token))
		token = util.EncryptAes(token)
		beego.Debug(fmt.Sprintf("cookie token: %s", token))
		c.Ctx.SetCookie("token", token, 0)
		c.JSON(base.SUCCESS)
	}
}

/**
退出登陆
*/
func (c *LoginController) LoginOut() {
	c.Ctx.SetCookie("token", "", 0)
	c.NewRedirect(beego.URLFor("LoginController.LoginView"))
}
