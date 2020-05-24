package repository

import (
	"main/db"
	"main/modules/model"
)

type UserRepository struct {

}

/**
保存用户信息
 */
func (user *UserRepository) SaveUser(userData *model.User) error {
	return db.G_DB.Model(&model.User{}).Save(&userData).Error
}

/**
AllUsers查询 用户信息
 */
func (user *UserRepository) AllUsers() ([]*model.User, error)  {
	var (
		err error
		result = make([]*model.User, 0)
	)
	if err = db.G_DB.
		Find(&result, "").Error; err != nil {
		return nil, err
	}
	return result, nil
}


