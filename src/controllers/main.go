package controllers

import (
	"github.com/go-steven/cms/src/controllers/base"
	"github.com/go-steven/cms/src/service"
)

type MainController struct {
	base.BaseController
}

/**
进入首页
*/
func (c *MainController) Index() {
	c.Show("index.html")
}

/**
欢迎页面
*/
func (c *MainController) Welcome() {
	c.Show("common/welcome.html")
}

/**
左侧菜单
*/
func (c *MainController) LeftMenu() {
	c.Show("common/leftMenu.html")
}

/**
头页面
*/
func (c *MainController) Header() {
	c.Data["user"] = c.User
	c.Show("common/header.html")
}

/**
进入没有权限页面
*/
func (c *MainController) NoRole() {
	c.Show("common/noRole.html")
}

/**
加载主页面权限tree
*/
func (c *MainController) LoadMenu() {
	roles := service.RoleService.LoadMenu(c.User.Id)
	c.JSON(roles)
}
