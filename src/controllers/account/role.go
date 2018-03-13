package account

import (
	"github.com/go-steven/cms/src/controllers/base"
	model_account "github.com/go-steven/cms/src/model/account"
	"github.com/go-steven/cms/src/service"
	"github.com/go-steven/cms/src/util"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type RoleController struct {
	base.BaseController
}

/**
进入分页展示页面
*/
func (c *RoleController) ListView() {
	c.Show("role/list.html")
}

/**
进入添加权限页面
*/
func (c *RoleController) AddView() {
	c.Show("role/add.html")
}

/**
进入添加权限目录页面
*/
func (c *RoleController) AddDirView() {
	c.Show("role/addDir.html")
}

/**
进入修改页面，根据ID查询权限对象
*/
func (c *RoleController) UpdateView() {
	id, _ := c.GetUint64("role_id")
	role, err := service.RoleService.GetRoleById(id)
	if err != nil {
		c.JSON(err.Error())
	}
	//c.JSON(role)
	c.Data["role"] = role
	c.Show("role/update.html")
}

/**
获取分页展示数据
*/
func (c *RoleController) GridList() {
	page, _ := c.GetInt("page")
	pagesize, _ := c.GetInt("rows")
	p := util.NewPager(page, pagesize)
	roleid, _ := c.GetInt("role_id")

	roleName := c.GetString("role_name")
	roleUrl := c.GetString("role_url")

	count, roles := service.RoleService.GridList(p, roleid, roleName, roleUrl)
	c.JSONPage(count, roles)
}

/**
加载权限树
*/
func (c *RoleController) ListTree() {
	id, _ := c.GetUint64("id")
	roles := service.RoleService.ListTree(true)
	//展开一级目录和当前添加节点的父节点（权限菜单一般只会有两级所以这样可以让当前添加的节点及时的展示出来）
	for i, role := range roles {
		if role.Pid == 0 {
			roles[i].Open = true
		}
		if role.Id == id {
			roles[i].Open = true
		}
	}
	c.JSON(roles)
}

/**
添加权限
*/
func (c *RoleController) AddRole() {
	pid, _ := c.GetUint64("pid")
	name := c.GetString("name")
	roleUrl := c.GetString("role_url")
	isMenu, _ := c.GetUint8("is_menu")
	remarks := c.GetString("remarks")
	module := c.GetString("module")
	action := c.GetString("action")

	//参数校验
	valid := validation.Validation{}
	valid.Required(name, "权限名称").Message("不能为空")
	valid.MaxSize(name, 20, "权限名称").Message("长度不能超过20个字符")
	valid.Required(remarks, "描述信息").Message("不能为空")
	valid.MaxSize(remarks, 50, "描述信息").Message("长度不能超过50个字符")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			c.JSON((err.Key + err.Message))
		}
	}

	role := &model_account.Role{
		Pid:     pid,
		Name:    name,
		RoleUrl: roleUrl,
		IsMenu:  isMenu,
		Remarks: remarks,
		Module:  module,
		Action:  action,
	}
	beego.Debug("add role:", role)

	if err := service.RoleService.AddRole(role); err != nil {
		c.JSON(err.Error())
	}
	c.JSON(base.SUCCESS)
}

/**
修改权限
*/
func (this *RoleController) Modify() {
	id, _ := this.GetUint64("id")
	pid, _ := this.GetUint64("pid")
	name := this.GetString("name")
	roleurl := this.GetString("role_url")
	isMenu, _ := this.GetUint8("is_menu")
	remarks := this.GetString("remarks")
	module := this.GetString("module")
	action := this.GetString("action")

	//参数校验
	valid := validation.Validation{}
	valid.Required(name, "权限名称").Message("不能为空")
	valid.MaxSize(name, 20, "权限名称").Message("长度不能超过20个字符")
	valid.Required(remarks, "描述信息").Message("不能为空")
	valid.MaxSize(remarks, 50, "描述信息").Message("长度不能超过50个字符")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			this.JSON((err.Key + err.Message))
		}
	}

	role := &model_account.Role{
		Id:      id,
		Pid:     pid,
		Name:    name,
		RoleUrl: roleurl,
		IsMenu:  isMenu,
		Remarks: remarks,
		Module:  module,
		Action:  action,
	}
	beego.Debug(role)
	if err := service.RoleService.ModifyRole(role); err != nil {
		this.JSON("修改失败！")
	}
	this.JSON(base.SUCCESS)
}

/**
删除权限
*/
func (this *RoleController) DeleteRole() {
	ids := this.GetStrings("ids")

	if err := service.RoleService.DeleteRole(ids); err != nil {
		this.JSON(err.Error())
	}
	this.JSON(base.SUCCESS)
}
