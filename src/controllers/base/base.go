package base

import (
	"fmt"
	model_account "github.com/go-steven/cms/src/model/account"
	"github.com/go-steven/cms/src/service"
	"github.com/go-steven/cms/src/util"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

const (
	SUCCESS = "success"
)

type BaseController struct {
	beego.Controller

	User           *model_account.User // 当前登录的用户id
	ControllerName string              // 控制器名
	ActionName     string              // 动作名
	OpenPerm       map[string]bool     // 公开的权限
}

/*
登陆鉴权等操作，
测试开发的时候可以注释这个方法，方便测试
*/
func (c *BaseController) Prepare() {
	c.Ctx.Output.Header("Author", "Steven YANG")
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "")

	//获取请求方法名称
	controllerName, actionName := c.GetControllerAndAction()
	c.ControllerName = controllerName
	c.ActionName = actionName

	//判断是否是不需要鉴权的公共操作
	if c.IsOpenPerm() {
		return
	}

	//登录校验
	token := c.Ctx.GetCookie("token")
	if user := ValidateToken(token, c.GetClientIp()); user == nil {
		c.NewRedirect(beego.URLFor("LoginController.Tologin"))
	} else {
		c.User = user
	}

	// //TODO 暂时判断如果是admin账号登陆就不执行任何权限校验，后续改为在某个组的用户都不做校验
	// if strings.EqualFold(service.RoleService.IsAdministrator(c.user.Id)) {
	// 	return
	// }

	if strings.EqualFold(controllerName, "MainController") {
		return
	}

	//操作权限校验
	if ok, err := c.ValidateRole(); !ok {
		if c.IsAjax() {
			c.JSON(err.Error())
		} else {
			c.NewRedirect(beego.URLFor("MainController.Norole"))
		}
	}

}

/**
初始化开放权限(不需要权限校验的操作,后续如果有不需要权限校验的操作都可以写在这里)
*/
func (c *BaseController) InitOpenPerm() {
	c.OpenPerm = map[string]bool{
		"MainController.LeftMenu": true,
		"MainController.Norole":   true,
	}
}

/**
判断是否是不需要鉴权的公共操作
*/
func (c *BaseController) IsOpenPerm() bool {
	//如果是登陆相关操作则不进行登陆鉴权和权限鉴权等操作
	if strings.EqualFold(c.ControllerName, "logincontroller") {
		return true
	}
	c.InitOpenPerm()
	v, ok := c.OpenPerm[c.ControllerName+"."+c.ActionName]
	return ok && v
}

/**
token 校验，判断是否登录
*/
func ValidateToken(token, currentIp string) *model_account.User {
	dToken, err := util.DecryptAes(token)
	if err != nil {
		beego.Debug("token 解密失败")
		return nil
	}
	arr := strings.Split(dToken, "|")
	if len(arr) != 3 {
		beego.Debug("token 校验失败")
		return nil
	}
	userIdStr := arr[0]
	ip := arr[2]
	if !strings.EqualFold(ip, currentIp) {
		//IP发生变化 强制重新登录
		beego.Debug("ip changed")
		return nil
	}
	userId, _ := strconv.ParseUint(userIdStr, 10, 64)
	user, err := service.UserService.GetUserById(userId)
	if err != nil || user.Id < 0 {
		beego.Debug("ID error")
		return nil
	}
	return user
}

/**
校验权限
*/
func (c *BaseController) ValidateRole() (bool, error) {
	if err := service.RoleService.ValidateRole(c.ControllerName, c.ActionName, c.User.Id); err != nil {
		return false, err
	}
	return true, nil
}

/**
重定向
*/
func (c *BaseController) NewRedirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

/*
指定页面，并且返回公共参数
*/
func (c *BaseController) Show(url string) {
	c.Data["staticUrl"] = beego.AppConfig.String("staticUrl")
	c.TplName = url
}

/**
把需要返回的结构序列化成json 输出
*/
func (c *BaseController) JSON(result interface{}) {
	c.Data["json"] = result
	c.ServeJSON()
	c.StopRun()
}

type Empty struct {
}

/*
 用于分页展示列表的时候的 输出json
*/
func (c *BaseController) JSONPage(total int, rows interface{}) {
	beego.Debug("分页数据：", total, rows)
	ret := make(map[string]interface{}, 1)
	if total == 0 || rows == nil {
		beego.Debug("查询分页数据为空，返回默认json")
		//这里默认total设置为1是因为easyui分页控件如果total 为0会出现错乱
		ret["total"] = 1
		ret["rows"] = make([]Empty, 0)
	} else {
		ret["total"] = total
		ret["rows"] = rows
	}
	c.Data["json"] = ret
	c.ServeJSON()
	c.StopRun()
}

/**
获取IP
*/
func (c *BaseController) GetClientIp() string {
	ip := c.Ctx.Request.Header.Get("Remote_addr")
	if ip == "" {
		ip = c.Ctx.Request.RemoteAddr
	}
	fmt.Println(ip)
	if strings.Contains(ip, ":") {
		ip = util.SubString(ip, 0, strings.Index(ip, ":"))
	}
	fmt.Println(ip)
	return ip
}
