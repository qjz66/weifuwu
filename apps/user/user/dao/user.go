package dao

import "tiny_service/apps/user/user/internal/models"

func AddUser(user *models.User) error {
	err := models.UserMgr(gdb).Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByPwd(username, password string) (*models.User, error) {
	mgr := models.UserMgr(gdb)
	u, err := mgr.GetByOption(mgr.WithUsername(username), mgr.WithPassword(password))
	if err != nil {
		return nil, err
	}
	return &u, nil
}
