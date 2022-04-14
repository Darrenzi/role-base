package service

import (
	"blog/common/global"
	"blog/common/util"
	"blog/model"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserService struct {
}

func (u *UserService) Add(user *model.User, roles []string) (uint, error) {
	db := global.GetDB()
	var count int64
	db.Model(user).Where("username = ?", user.Username).Count(&count)
	if count != 0 {
		return 0, errors.New("用户名已存在")
	}

	user.Salt = util.GetRandomString(32)
	password := user.Password + user.Salt
	user.Password = util.GenerateMD5(password)

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := db.Create(user).Error; err != nil {
			return err
		}

		if _, err := u.Authorize(int(user.ID), roles); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Errorf("新增用户失败: ", err)
		return 0, err
	}

	return user.ID, nil
}

func (u *UserService) Delete(user *model.User) (bool, error) {
	db := global.GetDB()

	tx := db.Model(&user)

	if user.Username != "" {
		tx.Where("username=?", user.Username)
	}

	if err := tx.Delete(user).Error; err != nil {
		log.Error("删除用户失败: ", err)
		return false, err
	}

	return true, nil
}

func (u *UserService) Find(user *model.User) ([]model.User, error) {
	db := global.GetDB()

	var users []model.User
	if err := db.Preload("Roles").Where(user).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) Roles(id int) ([]model.Role, error) {
	db := global.GetDB()
	var roles []model.Role
	if err := db.Where("id in (?)",
		db.Table("user_role").Select("role_id").Where("user_id = ?", id)).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (u *UserService) Authorize(userID int, roles []string) (bool, error) {
	db := global.GetDB()
	for _, roleId := range roles {
		var count int64
		db.Table("user_role").Where("user_id=? AND role_id=?", userID, roleId).Count(&count)

		if count == 0 {
			if err := db.Table("user_role").Create(map[string]interface{}{
				"user_id": userID,
				"role_id": roleId,
			}).Error; err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

func (u *UserService) Login(form model.LoginForm) (string, error) {
	db := global.GetDB()
	var user model.User
	db.Select("id", "password", "salt").Where("username =?", form.Username).First(&user)

	password := util.GenerateMD5(form.Password + user.Salt)
	if password == user.Password {
		rows, err := db.Table("roles").Select("name").Where("id IN (?)", db.Table("user_role").Select("role_id").Where("user_id=?", user.ID)).Rows()
		if err != nil {
			log.Error("获取用户橘色失败: ", err)
			return "", err
		}

		var roles []string
		for rows.Next() {
			var role string
			rows.Scan(&role)
			roles = append(roles, role)
		}

		token, err := global.GenerateToken(int(user.ID), user.Username, roles)
		if err != nil {
			log.Error("生成token失败:", err)
			return "", err
		}
		return token, nil
	} else {
		return "", errors.New("账号或密码错误")
	}

}
