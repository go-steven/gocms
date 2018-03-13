package account

import (
	"github.com/go-steven/cms/src/controllers/base"
	model_account "github.com/go-steven/cms/src/model/account"
	"github.com/go-steven/cms/src/service"
	"github.com/go-steven/cms/src/util"
	"time"

	"github.com/astaxie/beego/validation"
)

type UserController struct {
	base.BaseController
}

/**
进入管理员列表页面
*/
func (c *UserController) ListView() {
	c.Show("user/list.html")
}

/**
进入添加页面
*/
func (c *UserController) AddView() {
	c.Show("user/add.html")
}

/**
进入修改页面
*/
func (c *UserController) UpdateView() {
	userId, _ := c.GetUint64("user_id")
	user, _ := service.UserService.GetUserById(userId)
	c.Data["user"] = user
	c.Show("user/update.html")
}

/**
获取分页展示数据
*/
func (c *UserController) GridList() {
	page, _ := c.GetInt("page")
	pagesize, _ := c.GetInt("rows")
	mail := c.GetString("mail")
	phone := c.GetString("phone")
	realname := c.GetString("realname")
	userId := c.GetString("id")
	name := c.GetString("name")
	p := util.NewPager(page, pagesize)
	total, rows := service.UserService.GridList(p, userId, mail, realname, phone, name)
	c.JSONPage(total, rows)
}

/**
添加管理员
*/
func (c *UserController) AddUser() {
	name := c.GetString("name")
	mail := c.GetString("mail")
	realname := c.GetString("realname")
	phone := c.GetString("phone")
	department := c.GetString("department")
	passwd := c.GetString("passwd")
	groupIds := c.GetString("group_ids")

	//参数校验
	valid := validation.Validation{}
	valid.Required(name, "账号").Message("不能为空")
	valid.MaxSize(name, 20, "账号").Message("长度不能超过20个字符")
	valid.Required(mail, "邮箱").Message("不能为空")
	valid.MaxSize(mail, 50, "邮箱").Message("长度不能超过50个字符")
	valid.Email(mail, "邮箱").Message("格式错误")
	valid.Required(realname, "姓名").Message("不能为空")
	valid.MaxSize(realname, 20, "姓名").Message("长度不能超过20个字符")
	valid.Required(phone, "手机号码").Message("不能为空")
	valid.MaxSize(phone, 15, "手机号码").Message("长度不能超过15个字符")
	valid.Required(department, "部门").Message("不能为空")
	valid.MaxSize(department, 20, "部门").Message("长度不能超过20个字符")
	valid.Required(passwd, "密码").Message("不能为空")
	valid.MaxSize(passwd, 20, "密码").Message("长度不能超过20个字符")
	valid.MinSize(groupIds, 1, "组信息").Message("请至少选择一个")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			c.JSON((err.Key + err.Message))
		}
	}

	passwd = util.EncodeMessageMd5(passwd)

	user := &model_account.User{
		Name:       name,
		Realname:   realname,
		Mail:       mail,
		Phone:      phone,
		Department: department,
		Passwd:     passwd,
		Created:    time.Now(),
		Updated:    time.Now(),
	}
	if err := service.UserService.AddUser(user, groupIds); err != nil {
		c.JSON(err.Error())
	}
	c.JSON(base.SUCCESS)
}

/**
修改管理员
*/
func (c *UserController) ModifyUser() {
	userId, _ := c.GetUint64("user_id")
	name := c.GetString("name")
	mail := c.GetString("mail")
	realname := c.GetString("realname")
	phone := c.GetString("phone")
	department := c.GetString("department")
	passwd := c.GetString("passwd")
	groupIds := c.GetString("group_ids")

	//参数校验
	valid := validation.Validation{}
	valid.Required(name, "账号").Message("不能为空")
	valid.MaxSize(name, 20, "账号").Message("长度不能超过20个字符")
	valid.Required(mail, "邮箱").Message("不能为空")
	valid.MaxSize(mail, 50, "邮箱").Message("长度不能超过50个字符")
	valid.Email(mail, "邮箱").Message("格式错误")
	valid.Required(realname, "姓名").Message("不能为空")
	valid.MaxSize(realname, 20, "姓名").Message("长度不能超过20个字符")
	valid.Required(phone, "手机号码").Message("不能为空")
	valid.MaxSize(phone, 15, "手机号码").Message("长度不能超过15个字符")
	valid.Required(department, "部门").Message("不能为空")
	valid.MaxSize(department, 20, "部门").Message("长度不能超过20个字符")

	if len(passwd) > 0 {
		valid.Required(passwd, "密码").Message("不能为空")
		valid.MaxSize(passwd, 20, "密码").Message("长度不能超过20个字符")
	}

	valid.MinSize(groupIds, 1, "组信息").Message("请至少选择一个")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			c.JSON((err.Key + err.Message))
		}
	}

	if len(passwd) != 0 {
		passwd = util.EncodeMessageMd5(passwd)
	}

	user := &model_account.User{
		Id:         userId,
		Name:       name,
		Realname:   realname,
		Mail:       mail,
		Phone:      phone,
		Department: department,
		Passwd:     passwd,
		Created:    time.Now(),
		Updated:    time.Now(),
		IsDel:      0,
	}

	if err := service.UserService.ModifyUser(user, groupIds); err != nil {
		c.JSON(err.Error())
	}
	c.JSON(base.SUCCESS)
}

/**
删除
*/
func (this *UserController) Delete() {
	userids := this.GetString("user_ids")
	if err := service.UserService.Delete(userids); err != nil {
		this.JSON(err.Error())
	}
	this.JSON(base.SUCCESS)
}

/**
获取管理员组列表数据
修改管理员的时候需要加载管理员组列表，并且设置已经选择的权限为选中状态
*/
func (c *UserController) GridGroupList() {
	userId, _ := c.GetUint64("user_id")
	groupName := c.GetString("group_name")
	page, _ := c.GetInt("page")
	pagesize, _ := c.GetInt("rows")
	p := util.NewPager(page, pagesize)

	count, userGroup := service.UserGroupService.GridList(groupName, p)
	checkedGroupId := service.UserService.GetAllCheckGroup(userId)

	userCheckGroup := make([]model_account.UserGroupCheck, len(userGroup))

	for index, user := range userGroup {
		userCheck := model_account.UserGroupCheck{
			Id:        user.Id,
			GroupName: user.GroupName,
			Remarks:   user.Remarks,
			Created:   user.Created,
			Updated:   user.Updated,
			IsDel:     user.IsDel,
			Check:     checkedGroupId[user.Id],
		}
		userCheckGroup[index] = userCheck
	}

	c.JSONPage(count, userCheckGroup)
}
