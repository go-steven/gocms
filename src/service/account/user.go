package account

import (
	model_account "github.com/go-steven/cms/src/model/account"
	"github.com/go-steven/cms/src/service/base"
	"github.com/go-steven/cms/src/util"
	"strings"
	"time"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserService struct {
	base.BaseService
}

func NewUserService(o orm.Ormer) *UserService {
	return &UserService{
		base.BaseService{
			Ormer: o,
		},
	}
}

/**
分页查询管理员列表
*/
func (s *UserService) GridList(pager util.Pager, userid, usermail, realname, userphone, name string) (count int, users []model_account.User) {
	countsql := "SELECT COUNT(1) FROM t_user AS t "
	condition := genUserCondition(userid, usermail, realname, userphone, name)
	if err := s.Ormer.Raw(countsql + condition).QueryRow(&count); err != nil || count < 1 {
		beego.Debug("select user count err or result is null.")
		return
	}

	listsql := "SELECT id, name, mail, realname, phone, department, passwd, created, updated, is_del FROM t_user AS t "
	if _, err := s.Ormer.Raw(listsql+condition+util.LIMIT, pager.Offset, pager.PageSize).QueryRows(&users); err != nil {
		beego.Warn("select userList from db error.")
		return
	}
	return
}

/**
按照参数拼接sql查询条件
*/
func genUserCondition(userid, usermail, realname, userphone, name string) (condition string) {
	condition = " WHERE t.is_del = 0 "
	if !strings.EqualFold(userid, "") {
		condition += " AND t.id = " + userid + "'"
	}
	if !strings.EqualFold(usermail, "") {
		condition += " AND t.mail = '" + usermail + "'"
	}
	if !strings.EqualFold(realname, "") {
		condition += " AND t.realname =  '" + realname + "'"
	}
	if !strings.EqualFold(userphone, "") {
		condition += " AND t.phone =  '" + userphone + "'"
	}
	if !strings.EqualFold(name, "") {
		condition += " AND t.name =  '" + name + "'"
	}
	beego.Debug("condition is : ", condition)
	return
}

/**
添加管理员
*/
func (s *UserService) AddUser(user *model_account.User, groupIds string) error {
	flag := false
	if userId, err := s.Ormer.Insert(user); err != nil {
		beego.Warn("insert user fail, user:", user, err.Error())
		return &util.BizError{"添加失败,账号已经存在"}
	} else {
		idArray := strings.Split(groupIds, ",")
		for _, gid := range idArray {
			gidint, err := strconv.ParseUint(gid, 10, 64)
			if err != nil {
				beego.Debug("id 转换成数字异常，id：", gid)
				flag = true
			}
			rel := model_account.UserGroupRel{
				UserId:  uint64(userId),
				GroupId: gidint,
				IsDel:   0,
			}
			if _, err := s.Ormer.Insert(&rel); err != nil {
				flag = true
			}
		}
	}
	if flag {
		return &util.BizError{"出现异常，部分权限添加失败，请补充添加权限。"}
	}
	return nil
}

/**
修改管理员
*/
func (s *UserService) ModifyUser(user *model_account.User, groupIds string) error {
	flag := false
	updateSql := "UPDATE t_user SET "

	set := updateSet(user)
	condition := " WHERE id = ? "

	// if _, err := s.Ormer.Raw(updateSql, user.Name, user.Mail, user.RealName, user.Phone, user.Department, time.Now(), user.Id).Exec(); err != nil {
	id := user.Id
	if _, err := s.Ormer.Raw(updateSql+set+condition, id).Exec(); err != nil {
		beego.Warn("update user fail, user:", user, err.Error())
		return &util.BizError{"修改失败"}
	} else {
		//逻辑删除所有用户和组关联关系UserGroupRel
		delRelSql := "UPDATE t_user_group_rel SET is_del = 1 WHERE user_id = ?"
		if _, err := s.Ormer.Raw(delRelSql, user.Id).Exec(); err != nil {
			return &util.BizError{"修改失败"}
		}

		idArray := strings.Split(groupIds, ",")
		//重新添加关联关系
		for _, gid := range idArray {
			gidint, err := strconv.ParseUint(gid, 10, 64)
			if err != nil {
				beego.Debug("id 转换成数字异常，id：", gid)
				flag = true
			}
			rel := model_account.UserGroupRel{
				UserId:  user.Id,
				GroupId: gidint,
				IsDel:   0,
			}
			if _, err := s.Ormer.Insert(&rel); err != nil {
				beego.Warn("添加组关系失败", rel, err.Error())
				flag = true
			}
		}
	}
	if flag {
		return &util.BizError{"出现异常，部分权限修改失败，请补充添加权限。"}
	}

	return nil
}

func updateSet(user *model_account.User) string {
	set := ""
	if !strings.EqualFold(user.Passwd, "") {
		set += " passwd = '" + user.Passwd + "',"
	}
	if !strings.EqualFold(user.Name, "") {
		set += " name = '" + user.Name + "',"
	}
	if !strings.EqualFold(user.Mail, "") {
		set += " mail = '" + user.Mail + "',"
	}
	if !strings.EqualFold(user.Realname, "") {
		set += " realname = '" + user.Realname + "',"
	}
	if !strings.EqualFold(user.Phone, "") {
		set += " phone = '" + user.Phone + "',"
	}
	if !strings.EqualFold(user.Department, "") {
		set += " department = '" + user.Department + "',"
	}
	set += " updated = '" + time.Now().Format("2006-01-02 15:04:05") + "'"

	return set
}

/**
删除管理员基本信息
*/
func (s *UserService) Delete(userids string) error {
	delUserSql := "UPDATE t_user SET is_del = 1 WHERE id in (" + userids + ")"
	if _, err := s.Ormer.Raw(delUserSql).Exec(); err != nil {
		return &util.BizError{"删除管理员基本信息失败"}
	}
	delRelSql := "UPDATE t_user_group_rel SET is_del = 1 WHERE user_id IN (" + userids + ")"
	if _, err := s.Ormer.Raw(delRelSql).Exec(); err != nil {
		return &util.BizError{"删除管理员和组关系失败"}
	}
	return nil
}

/**
登陆鉴权
*/
func (s *UserService) Authentication(name, encodePwd string) (user *model_account.User, err error) {
	sql := "SELECT id, passwd FROM t_user AS t WHERE t.name = '" + name + "' AND is_del = 0"
	if err := s.Ormer.Raw(sql).QueryRow(&user); err != nil {
		if err == orm.ErrNoRows {
			return nil, &util.BizError{"账号不存在"}
		}
		return nil, &util.BizError{"登陆失败，请稍后重试"}
	}
	if !strings.EqualFold(encodePwd, user.Passwd) {
		return nil, &util.BizError{"密码错误"}
	}
	return user, nil
}

/**
根据ID查询管理员
*/
func (s *UserService) GetUserById(id uint64) (user *model_account.User, err error) {
	user = &model_account.User{Id: id}
	if err := s.Ormer.Read(user); err != nil {
		if err == orm.ErrNoRows {
			err = &util.BizError{"账号不存在"}
			return nil, err
		}
		err = &util.BizError{"系统错误"}
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAllCheckGroup(id uint64) map[uint64]bool {
	var list orm.ParamsList
	num, err := s.Ormer.Raw("SELECT group_id FROM t_user_group_rel AS t WHERE is_del = 0 AND t.user_id = ?", id).ValuesFlat(&list)
	if err != nil || num < 1 {
		return nil
	}
	roleIdMap := make(map[uint64]bool, len(list))
	for i := 0; i < len(list); i++ {
		idStr := list[i].(string)
		id, _ := strconv.ParseUint(idStr, 10, 64)
		roleIdMap[id] = true
	}
	return roleIdMap
}
