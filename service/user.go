package service

import (
	"errors"
	"fmt"
	"go-blog/global"
	"go-blog/model"
	"go-blog/model/request"
	"go-blog/model/scopes"
	"go-blog/utils"
	"gorm.io/gorm"
)

// 查询用户是否存在
func CheckUserExist(user *model.User) bool {
	u, err := FindUser(&model.User{Base: model.Base{ID: user.ID}, Username: user.Username})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return u.ID != ""
}

func Login(login *request.LoginBody) (*model.User, string, error) {
	user, err := FindUser(&model.User{Username: login.Username, Password: login.Password})
	if err != nil {
		return nil, "", err
	}

	if user.ID == "" {
		return nil, "", nil
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	return user, token, err
}

func CreateUser(d *request.CreateUserBody) error {
	var err error
	user := &model.User{}

	if err = utils.FieldValTransfer(user, d); err != nil {
		return err
	}

	if err = utils.FieldValTransfer(&user.UserInfo, d); err != nil {
		return err
	}

	if d.Password != "" {
		user.Password, err = utils.EncryptSha256(d.Password)
		if err != nil {
			return err
		}
	}

	if user.Nickname == "" {
		user.Nickname = user.Username
	}

	if user.UserInfo.Description == "" {
		user.UserInfo.Description = "这家伙很懒，啥也不写。"
	}

	if len(d.Roles) > 0 {
		roles := []model.Role{}
		for _, id := range d.Roles {
			roles = append(roles, model.Role{Base: model.Base{ID: id}})
		}
		user.Roles = roles
	}

	return global.DB.Create(user).Error
}

func FindSelfInfo(token string) (*model.User, error) {
	user := &model.User{}
	u, err := utils.ParseToken(token)
	if err != nil {
		return nil, err
	}

	user.ID = u.ID
	err = global.DB.
		Preload("UserInfo").
		Preload("Roles").
		First(user).Error

	return user, err
}

func FindUser(user *model.User) (*model.User, error) {
	u := &model.User{}

	err := global.DB.
		Preload("UserInfo").
		Preload("Roles", func(db *gorm.DB) *gorm.DB {
			return db.Select("id,name")
		}).Where(user).First(u).Error
	if err != nil {
		return nil, err
	}

	return u, err
}

func GetUserPaging(params *request.FindUserBody, pageInfo *request.Pagination, count *int64) ([]model.User, error) {
	var err error
	res := []model.User{}

	db := global.DB.Model(&model.User{}).
		Where(fmt.Sprintf("username like  '%%%s%%' and phone like '%%%s%%' and email like '%%%s%%'", params.Username, params.Phone, params.Email))

	if pageInfo != nil {
		db = db.Scopes(scopes.Pager(pageInfo.Page, pageInfo.PageSize, count))
	}

	err = db.Scan(&res).Error
	if err != nil {
		return []model.User{}, err
	}

	return res, err
}

func UpdateUser(params *request.UpdateUserBody) error {
	var err error
	user := model.User{}
	userInfo := user.UserInfo
	roles := []model.Role{}

	if err = utils.FieldValTransfer(&user, params); err != nil {
		return err
	}

	if err = utils.FieldValTransfer(&userInfo, params); err != nil {
		return err
	}

	db := global.DB.Model(&model.User{Base: model.Base{ID: params.ID}})
	if params.Roles != nil {
		l := len(params.Roles)
		if l > 0 {
			for i := 0; i < l; i++ {
				roles = append(roles, model.Role{Base: model.Base{ID: params.Roles[i]}})
			}
			err = db.Association("Roles").Replace(roles)
		} else {
			err = db.Association("Roles").Clear()
		}
	}

	if err != nil {
		return err
	}

	err = global.DB.Where("id = ?", params.ID).Updates(&user).Error
	if err != nil {
		return err
	}

	return global.DB.Table("user_info").Where("user_id = ?", params.ID).Updates(&userInfo).Error
}

func DeleteUser(ids []string) error {
	return global.DB.Delete(&model.User{}, ids).Error
}

func PatchAddRole(userIds []string, roleIds []string) error {
	users := []model.User{}
	roles := []model.Role{}

	for i := 0; i < len(userIds); i++ {
		users = append(users, model.User{Base: model.Base{ID: userIds[i]}})
	}

	for i := 0; i < len(roleIds); i++ {
		roles = append(roles, model.Role{Base: model.Base{ID: roleIds[i]}})
	}

	return global.DB.Model(&users).Association("Roles").Append(roles)
}

func PatchRemoveRole(userIds []string, roleIds []string) error {
	users := []model.User{}

	for i := 0; i < len(userIds); i++ {
		users = append(users, model.User{Base: model.Base{ID: userIds[i]}})
	}

	return global.DB.Model(&users).Where("role_id in (?)", roleIds).Association("Roles").Clear()
}
