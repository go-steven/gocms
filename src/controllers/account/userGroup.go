package account

import (
	"github.com/go-steven/cms/src/controllers/base"
	model_account "github.com/go-steven/cms/src/model/account"
	"github.com/go-steven/cms/src/service"
	"github.com/go-steven/cms/src/util"
	"time"

	"github.com/astaxie/beego/validation"
)

type UserGroupController struct {
	base.BaseController
}

/**
进入管理员组管理页面
*/
func (c *UserGroupController) ListView() {
	c.Show("usergroup/list.html")
}

/**
进入添加页面
*/
func (c *UserGroupController) AddView() {
	c.Show("usergroup/add.html")
}

/**
进入修改管理员组页面
*/
func (c *UserGroupController) UpdateView() {
	id, _ := c.GetUint64("user_group_id")
	usergroup := service.UserGroupService.GetUserGroupById(id)
	c.Data["user_group"] = usergroup
	c.Show("usergroup/update.html")
}

/**
获取管理员组列表数据
*/
func (c *UserGroupController) GridList() {
	groupName := c.GetString("group_name")
	page, _ := c.GetInt("page")
	pagesize, _ := c.GetInt("rows")
	p := util.NewPager(page, pagesize)

	count, userGroup := service.UserGroupService.GridList(groupName, p)
	c.JSONPage(count, userGroup)
}

/**
添加管理员组
*/
func (c *UserGroupController) AddUserGroup() {
	ids := c.GetString("ids")
	groupname := c.GetString("group_name")
	remarks := c.GetString("remarks")

	//参数校验
	valid := validation.Validation{}
	valid.Required(groupname, "管理员组名称").Message("不能为空")
	valid.MaxSize(groupname, 20, "管理员组名称").Message("长度不能超过20个字符")
	valid.Required(remarks, "描述信息").Message("不能为空")
	valid.MaxSize(remarks, 50, "描述信息").Message("长度不能超过50个字符")
	valid.MinSize(ids, 1, "权限").Message("请至少选择一个")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			c.JSON((err.Key + err.Message))
		}
	}

	usergroup := &model_account.UserGroup{
		GroupName: groupname,
		Remarks:   remarks,
		Created:   time.Now(),
		Updated:   time.Now(),
	}
	if err := service.UserGroupService.AddUserGroup(usergroup, ids); err != nil {
		c.JSON(err.Error())
	}
	c.JSON(base.SUCCESS)
}

/**
修改管理员组
*/
func (c *UserGroupController) ModifyUserGroup() {
	ids := c.GetString("ids")
	groupname := c.GetString("group_name")
	remarks := c.GetString("remarks")
	id, _ := c.GetUint64("id")

	//参数校验
	valid := validation.Validation{}
	valid.Required(groupname, "管理员组名称").Message("不能为空")
	valid.MaxSize(groupname, 20, "管理员组名称").Message("长度不能超过20个字符")
	valid.Required(remarks, "描述信息").Message("不能为空")
	valid.MaxSize(remarks, 50, "描述信息").Message("长度不能超过50个字符")
	valid.MinSize(ids, 1, "权限").Message("请至少选择一个")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			c.JSON((err.Key + err.Message))
		}
	}

	usergroup := &model_account.UserGroup{
		Id:        id,
		GroupName: groupname,
		Remarks:   remarks,
		Created:   time.Now(),
		Updated:   time.Now(),
	}
	if err := service.UserGroupService.ModifyUserGroup(usergroup, ids); err != nil {
		c.JSON(err.Error())
	}
	c.JSON(base.SUCCESS)
}

/**
删除管理员组
*/
func (c *UserGroupController) Delete() {
	ids := c.GetString("ids")
	if err := service.UserGroupService.Delete(ids); err != nil {
		c.JSON(err.Error())
	}
	c.JSON(base.SUCCESS)
}

/**
加载权限树(用于添加管理员组的时候选择权限)
*/
func (c *UserGroupController) LoadTreeWithoutRoot() {
	//查询树结构不加载root节点
	roles := service.RoleService.ListTree(false)
	//展开一级目录
	for i, role := range roles {
		if role.Pid == 0 {
			roles[i].Open = true
		}
	}
	c.JSON(roles)
}

/**
加载权限树(用于修改管理员组的时候选择权限-添加时选择的权限在修改的时候需要选中)
*/
func (c *UserGroupController) LoadTreeChecked() {
	groupUserId, _ := c.GetUint64("group_user_id")
	roleIdMap := service.UserGroupService.GetAllRoleByGroupId(groupUserId)
	//查询树结构不加载root节点
	roles := service.RoleService.ListTree(false)
	if roleIdMap == nil {
		//展开一级目录
		for i, role := range roles {
			if role.Pid == 0 {
				roles[i].Open = true
			}
		}
	} else {
		for i, role := range roles {
			if role.Pid == 0 {
				roles[i].Open = true
			}
			if _, ok := roleIdMap[role.Id]; ok {
				roles[i].Checked = true
			}
		}
	}
	c.JSON(roles)
}
