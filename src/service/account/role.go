package account

import (
	"bytes"
	model_account "github.com/go-steven/cms/src/model/account"
	"github.com/go-steven/cms/src/service/base"
	"github.com/go-steven/cms/src/util"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RoleService struct {
	base.BaseService
}

func NewRoleService(o orm.Ormer) *RoleService {
	return &RoleService{
		base.BaseService{
			Ormer: o,
		},
	}
}

/**
添加权限
*/
func (s *RoleService) AddRole(role *model_account.Role) error {
	if _, err := s.Ormer.Insert(role); err != nil {
		return &util.BizError{"添加失败"}
	}
	return nil
}

/**
查询列表的分页数据
*/
func (s *RoleService) GridList(pager util.Pager, roleid int, roleName, roleUrl string) (int, []model_account.Role) {
	//查询总数
	contsql := "SELECT COUNT(1) FROM t_role AS t WHERE t.pid = ?"
	condition := genCondition(roleName, roleUrl)
	var count int
	err := s.Ormer.Raw(contsql+condition, roleid).QueryRow(&count)
	if err != nil {
		beego.Error("查询Pid为", roleid, "的role总数异常，error message：", err.Error())
	}
	beego.Debug("pid 为", roleid, "的role有", count, "个")

	if count < 1 {
		beego.Info("没有pid 为", roleid, "的role")
		return 0, nil
	}

	// 从数据库查询数据
	var roles []model_account.Role
	listsql := "SELECT id, pid, name, role_url, module, action, is_menu, remarks FROM t_role AS t WHERE t.pid = ?  "
	_, err = s.Ormer.Raw(listsql+condition+util.LIMIT, roleid, pager.Offset, pager.PageSize).QueryRows(&roles)
	if err != nil {
		beego.Error("查询Pid为", roleid, "的role列表异常，error message：", err.Error())
	}

	return count, roles
}

func genCondition(roleName, roleUrl string) (condition string) {
	if !strings.EqualFold(roleName, "") {
		condition += " AND t.name = '" + roleName + "'"
	}
	if !strings.EqualFold(roleUrl, "") {
		condition += " AND t.role_url = '" + roleUrl + "'"
	}
	return
}

/**
查询树
@param needRoot:查询的数据集中是否需要包含root节点
*/
func (s *RoleService) ListTree(needRoot bool) []model_account.RoleTree {
	var buf bytes.Buffer
	buf.WriteString("SELECT id, pid, name, role_url, is_menu, remarks FROM t_role AS t ")
	if !needRoot {
		buf.WriteString(" WHERE t.id != 0")
	}
	var roles []model_account.RoleTree
	beego.Debug("查询权限树sql：", buf.String())
	_, err := s.Ormer.Raw(buf.String()).QueryRows(&roles)
	if err != nil {
		beego.Error("查询权限树的role列表异常，error message：", err.Error())
	}
	beego.Debug("生成权限树的数据：", roles)
	return roles
}

/**
根据ID查询role
*/
func (s *RoleService) GetRoleById(id uint64) (model_account.Role, error) {
	role := model_account.Role{Id: id}
	if err := s.Ormer.Read(&role); err != nil {
		return model_account.Role{}, err
	}
	return role, nil
}

/**
修改权限
*/
func (s *RoleService) ModifyRole(r *model_account.Role) error {
	role := model_account.Role{Id: r.Id}
	//根据ID读取
	if err := s.Ormer.Read(&role); err != nil {
		return err
	}
	//修改
	if num, err := s.Ormer.Update(r); num <= 0 && err != nil {
		return err
	}
	return nil
}

/**
删除权限
*/
func (s *RoleService) DeleteRole(ids []string) error {
	idstr := strings.Join(ids, ",")

	var count int
	countSubRoleSql := "SELECT COUNT(1) FROM t_role WHERE pid IN (" + idstr + ")"
	s.Ormer.Raw(countSubRoleSql).QueryRow(&count)
	if count > 0 {
		return &util.BizError{"不能删除有子节点的权限，请先删除所有子节点！"}
	}

	sql := "DELETE FROM t_role WHERE id IN (" + idstr + ")"
	if _, err := s.Ormer.Raw(sql).Exec(); err != nil {
		return &util.BizError{"删除失败！"}
	}
	return nil
}

/**
权限校验
*/
func (s *RoleService) ValidateRole(controllerName, actionName string, id uint64) error {
	if s.isAdministrator(id) {
		beego.Debug("用户属于超级管理员，不用校验权限")
		return nil
	}
	selectSql := `SELECT 
		COUNT(1) 
FROM t_user_group_rel AS ur, 
     t_role AS r, 
     t_group_role_rel AS gr 
WHERE r.module = ? 
	AND r.action = ? 
	AND ur.user_id = ? 
	AND ur.group_id = gr.group_id 
	AND r.id = gr.role_id 
	AND ur.is_del = 0 
	AND gr.is_del = 0`
	var count int
	s.Ormer.Raw(selectSql, controllerName, actionName, id).QueryRow(&count)
	if count > 0 {
		return nil
	}
	return &util.BizError{"您没有权限执行此操作，请联系系统管理员。"}
}

/**
加载权限树
*/
func (s *RoleService) LoadMenu(userId uint64) []model_account.RoleTree {
	var roles []model_account.RoleTree
	if s.isAdministrator(userId) {
		sql := "SELECT t.id, pid, name, role_url, is_menu, remarks FROM t_role AS t WHERE t.id != 0 AND t.is_menu = 0"
		if _, err := s.Ormer.Raw(sql).QueryRows(&roles); err != nil {
			beego.Error("查询权限树的role列表异常，error message：", err.Error())
			return roles
		}
	} else {
		sql := `SELECT DISTINCT t.id, pid, name, role_url, is_menu, remarks 
FROM t_role AS t
INNER JOIN t_group_role_rel AS gr ON (
	gr.role_id = t.id AND gr.is_del = 0
)
INNER JOIN t_user_group_rel AS ug ON (
	ug.group_id = gr.group_id AND ug.is_del = 0
)
WHERE t.id != 0 AND t.is_menu = 0
	AND ug.user_id = ?`
		if _, err := s.Ormer.Raw(sql, userId).QueryRows(&roles); err != nil {
			beego.Error("查询权限树的role列表异常，error message：", err.Error())
			return roles
		}
	}

	pidMap := make(map[uint64]bool, 10)
	for _, role := range roles {
		pidMap[role.Pid] = true
	}

	for i, role := range roles {
		//展开所有父节点
		if pidMap[role.Id] {
			roles[i].Open = true
			continue
		}
		if !strings.EqualFold(role.RoleUrl, "") {
			click := "click: addTab('" + roles[i].Name + "','" + roles[i].RoleUrl + "')"
			roles[i].Click = click
		}
	}

	return roles
}

/*
判断当前用户是否属于 超级管理员
*/
func (s *RoleService) isAdministrator(userId uint64) bool {
	var (
		flag bool
		list orm.ParamsList
	)
	num, err := s.Ormer.Raw("SELECT group_id FROM t_user_group_rel AS t WHERE t.user_id = ? AND t.is_del = 0", userId).ValuesFlat(&list)
	if err != nil || num < 1 {
		return flag
	}
	for i := 0; i < len(list); i++ {
		groupId := list[i].(string)
		if v, err := strconv.ParseUint(groupId, 10, 64); err == nil {
			if v == 1 {
				return true
			}
		}
	}
	return flag
}
