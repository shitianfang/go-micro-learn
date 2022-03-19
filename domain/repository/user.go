package repository

import (
	"user/domain/model"

	"github.com/jinzhu/gorm"
)

// 定义接口
type IUserRepository interface {
	InitTable() error
	FindUserByName(string) (*model.User, error)
	FindUserByID(int64) (*model.User, error)
	CreateUser(*model.User) (int64, error)
	DeleteUserByID(int64) error
	UpdateUser(*model.User) error
	FindAll() ([]model.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository { // 返回一个遵循结构类型的结构体
	return &UserRepository{db: db} // 报错说明还没有完全实现接口方法
}

// 遵循结构的结构体
type UserRepository struct {
	db *gorm.DB
}

// 为UserRepository结构体增加方法
func (u *UserRepository) InitTable() error {
	return u.db.CreateTable(&model.User{}).Error //使用User结构体的实例化，创建表
}

func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{} //创建一个User结构体，下一步填充数据到这个结构体
	return user, u.db.Where("user_name = ?", name).Find(user).Error
}

func (u *UserRepository) FindUserByID(userID int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.db.First(user, userID).Error
}

func (u *UserRepository) CreateUser(user *model.User) (userID int64, err error) {
	return user.ID, u.db.Create(user).Error
}

func (u *UserRepository) DeleteUserByID(userID int64) error {
	return u.db.Where("id=?", userID).Delete(&model.User{}).Error
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.db.Model(user).Update(&user).Error
}

func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.db.Find(&userAll).Error
}
