package routers

import (
	"github.com/go-steven/cms/src/controllers"
	controller_account "github.com/go-steven/cms/src/controllers/account"

	"github.com/astaxie/beego"
)

func init() {
	beego.Info("Init routers start ...")

	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/welcome", &controllers.MainController{}, "*:Welcome")
	beego.Router("/leftMenu", &controllers.MainController{}, "*:LeftMenu")
	beego.Router("/header", &controllers.MainController{}, "*:Header")
	beego.Router("/loadMenu", &controllers.MainController{}, "*:LoadMenu")
	// beego.Router("/login", &controller_main.MainController{}, "*:Login")
	// beego.Router("/tologin", &controller_main.MainController{}, "*:Tologin")
	// beego.Router("/loginpage", &controller_main.MainController{}, "*:Loginpage")
	//自动绑定映射关系
	beego.AutoRouter(&controller_account.UserController{})
	beego.AutoRouter(&controller_account.RoleController{})
	beego.AutoRouter(&controller_account.UserGroupController{})
	beego.AutoRouter(&controllers.LoginController{})

	beego.Info("Init routers end.")
}
