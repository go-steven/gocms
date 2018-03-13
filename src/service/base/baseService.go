package base

import (
	"github.com/astaxie/beego/orm"
)

type BaseService struct {
	Ormer orm.Ormer
}
