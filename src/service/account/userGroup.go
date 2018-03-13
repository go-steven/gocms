package account

import (
	model_account "github.com/go-steven/cms/src/model/account"
	"github.com/go-steven/cms/src/service/base"
	"github.com/go-steven/cms/src/util"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserGroupService struct {
	base.BaseService
}

func NewUserGroupService(o orm.Ormer) *UserGroupService {
	return &UserGroupService{
		base.BaseService{
			Ormer: o,
		},
	}
}

/**
查询管理员组分页列表
*/
func (s *UserGroupService) GridList(groupName string, pager util.Pager) (count int, userGroup []model_account.UserGroup) {
	coutsql := "SELECT COUNT(1) FROM t_user_group AS t "
	condition := genUserGroupCondition(groupName)
	if err := s.Ormer.Raw(coutsql + condition).QueryRow(&count); err != nil || count < 1 {
		//如果查询出错或者查询结果为空返回默认空值
		return
	}

	listsql := "SELECT id, group_name, remarks, created, updated, is_del FROM t_user_group AS t "
	if num, err := s.Ormer.Raw(listsql+condition+util.LIMIT, pager.Offset, pager.PageSize).QueryRows(&userGroup); err != nil || num < 1 {
		//如果查询出错返回默认空值
		return
	}
	return
}

func genUserGroupCondition(groupName string) (condition string) {
	condition = " WHERE t.is_del = 0 "
	if groupName != "" {
		condition += " AND t.group_name = " + groupName
	}
	return
}

/**
添加管理员组
*/
func (s *UserGroupService) AddUserGroup(usergroup *model_account.UserGroup, ids string) error {
	id, err := s.Ormer.Insert(usergroup)
	if err != nil || id < 1 {
		return &util.BizError{"添加失败"}
	}
	flag := false
	idArray := strings.Split(ids, ",")
	for _, v := range idArray {
		beego.Debug("给ID为", id, "的管理员组添加", v, "权限")
		roleId, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			beego.Warn(v, "不是数字")
			flag = true
			continue
		}
		groupRoleRel := &model_account.GroupRoleRel{
			GroupId: uint64(id),
			RoleId:  roleId,
			IsDel:   0,
		}
		if _, err := s.Ormer.Insert(groupRoleRel); err != nil {
			beego.Warn("给ID为", id, "的管理员组添加", groupRoleRel.RoleId, "权限失败")
			flag = true
			continue
		}
	}
	if flag {
		return &util.BizError{"出现异常，部分权限添加失败，请补充添加权限。"}
	}
	return nil
}

/**
修改管理员组
*/
func (s *UserGroupService) ModifyUserGroup(usergroup *model_account.UserGroup, ids string) error {
	//修改基础信息
	if _, err := s.Ormer.Update(usergroup); err != nil {
		beego.Warn("update usergroup db error.", err.Error())
		return &util.BizError{"修改失败"}
	}

	id := usergroup.Id
	//删除当前组关联的所有权限
	delsql := "UPDATE t_group_role_rel AS t SET t.is_del = 0 WHERE t.group_id = ? AND t.is_del = 0"
	if _, err := s.Ormer.Raw(delsql, id).Exec(); err != nil {
		beego.Warn("del group's role fail.", err.Error())
		return &util.BizError{"修改失败"}
	}

	//重新添加权限
	flag := false
	idArray := strings.Split(ids, ",")
	for _, v := range idArray {
		beego.Debug("给ID为", id, "的管理员组添加", v, "权限")
		roleId, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			beego.Warn(v, "不是数字")
			flag = true
			continue
		}
		groupRoleRel := &model_account.GroupRoleRel{
			GroupId: id,
			RoleId:  roleId,
			IsDel:   0,
		}
		if _, err := s.Ormer.Insert(groupRoleRel); err != nil {
			beego.Warn("给ID为", id, "的管理员组添加", groupRoleRel.RoleId, "权限失败", err.Error())
			flag = true
			continue
		}
	}
	if flag {
		return &util.BizError{"出现异常，部分权限添加失败，请补充添加权限。"}
	}
	return nil
}

/**
删除管理员组
*/
func (s *UserGroupService) Delete(ids string) error {
	delsql := "UPDATE t_user_group AS t SET t.is_del = 1 WHERE t.id IN (" + ids + ")"
	if _, err := s.Ormer.Raw(delsql).Exec(); err != nil {
		beego.Warn("delete fail id:", ids, err.Error())
		return &util.BizError{"删除失败"}
	}

	//删除当前组关联的所有权限
	delrolesql := "UPDATE t_group_role_rel t SET t.is_del = 1 WHERE t.group_id IN (" + ids + ") AND t.is_del = 0"
	if _, err := s.Ormer.Raw(delrolesql).Exec(); err != nil {
		beego.Warn("del group's role fail.", err.Error())
		return &util.BizError{"删除失败"}
	}
	return nil
}

/**
根据ID获取管理员组信息
*/
func (s *UserGroupService) GetUserGroupById(id uint64) model_account.UserGroup {
	usergroup := model_account.UserGroup{Id: id}
	if err := s.Ormer.Read(&usergroup); err != nil {
		return model_account.UserGroup{}
	}
	return usergroup
}

/**
根据管理员组ID获取所有的权限列表
*/
func (s *UserGroupService) GetAllRoleByGroupId(id uint64) map[uint64]bool {
	var list orm.ParamsList
	num, err := s.Ormer.Raw("SELECT role_id FROM t_group_role_rel AS t WHERE t.group_id = ? AND t.is_del = 0", id).ValuesFlat(&list)
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
