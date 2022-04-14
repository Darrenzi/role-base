package service

import (
	"blog/common/global"
	"blog/model"
	"errors"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type PolicyService struct {
}

func (p *PolicyService) Add(policy *model.Policy) error {
	db := global.GetDB()

	var count int64
	db.Table("roles").Where("name=?", policy.Sub).Count(&count)
	if count == 0 {
		return errors.New("角色不存在")
	}

	gormAdapter, _ := gormadapter.NewAdapterByDB(db)
	enforcer, _ := casbin.NewEnforcer(global.Config.Casbin.Path, gormAdapter)
	_, err := enforcer.AddPolicy(policy.Sub, policy.Obj, policy.Method)
	if err != nil {
		return err
	}

	enforcer.SavePolicy()
	return nil
}

func (p *PolicyService) Delete(policy *model.Policy) error {
	db := global.GetDB()

	gormAdapter, _ := gormadapter.NewAdapterByDB(db)
	enforcer, _ := casbin.NewEnforcer(global.Config.Casbin.Path, gormAdapter)
	_, err := enforcer.RemovePolicy(policy.Sub, policy.Obj, policy.Method)
	if err != nil {
		return err
	}

	enforcer.SavePolicy()
	return nil
}
