package service

import (
	"blog/common/global"
	"blog/model"

	log "github.com/sirupsen/logrus"
)

type RoleService struct {
}

func (r *RoleService) SaveRole(role *model.Role) (uint, error) {
	db := global.GetDB()

	exist := model.Role{
		Name: role.Name,
	}

	db.Where(exist).First(&exist)
	if exist.ID != 0 {
		return exist.ID, nil
	}

	if err := db.Save(role).Error; err != nil {
		log.Error("新增用户失败: ", err)
		return 0, err
	}

	return role.ID, nil
}
