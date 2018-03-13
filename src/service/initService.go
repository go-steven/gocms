package service

import (
	model_account "github.com/go-steven/cms/src/model/account"
	service_account "github.com/go-steven/cms/src/service/account"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	o           orm.Ormer
	TablePrefix string // 表前缀

	RoleService      *service_account.RoleService
	UserGroupService *service_account.UserGroupService
	UserService      *service_account.UserService
)

func init() {
	beego.Info("init orm start...")
	TablePrefix = beego.AppConfig.String("db.prefix")

	dbType := beego.AppConfig.String("db_type")
	dsn := generateDSN()
	orm.RegisterDataBase("default", dbType, dsn)

	orm.RegisterModelWithPrefix(TablePrefix,
		new(model_account.Role),
		new(model_account.UserGroup),
		new(model_account.GroupRoleRel),
		new(model_account.User),
		new(model_account.UserGroupRel),
	)
	orm.RunSyncdb("default", false, true)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	o = orm.NewOrm()
	orm.RunCommand()

	beego.Info("init orm end.")
	//初始化service
	initService()
}

func generateDSN() string {
	dbHost := beego.AppConfig.String("db_host")
	dbPort := beego.AppConfig.String("db_port")
	dbUser := beego.AppConfig.String("db_user")
	dbPasswd := beego.AppConfig.String("db_pass")
	dbName := beego.AppConfig.String("db_name")

	//beego.Debug(dbHost, dbPort, dbUser, dbPasswd, dbName, dbType)
	// root:@tcp(127.0.0.1:3306)/test?charset=utf8
	dsn := dbUser + ":" + dbPasswd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"
	return dsn
}

func initService() {
	RoleService = service_account.NewRoleService(o)
	UserGroupService = service_account.NewUserGroupService(o)
	UserService = service_account.NewUserService(o)
}
