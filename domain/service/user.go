package service

import (
	"errors"
	"user/domain/model"
	"user/domain/repository"

	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(*model.User) (int64, error)
	DelUser(int64) error
	UpdUser(user *model.User) error
	FindUser(string) (*model.User, error)
	CheckPwd(userName string, pwd string) (ok bool, err error)
}

type UserDataService struct {
	UserRepository repository.UserRepository
}

func NewUserDataService(userRepository repository.UserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("password is error")
	}
	return true, nil
}

func (u *UserDataService) AddUser(user *model.User) (userID int64, err error) {
	pwdByte, err := GeneratePassword(user.PassWord)
	if err != nil {
		return user.ID, err
	}
	user.PassWord = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

func (u *UserDataService) DelUser(UserID int64) error {
	return u.UserRepository.DeleteUserByID(UserID)
}

func (u *UserDataService) UpdUser(user *model.User) error {
	pwdByte, err := GeneratePassword(user.PassWord)
	if err != nil {
		return err
	}
	user.PassWord = string(pwdByte)
	return u.UserRepository.UpdateUser(user)
}

func (u *UserDataService) FindUser(username string) (user *model.User, err error) {
	return u.UserRepository.FindUserByName(username)
}

func (u *UserDataService) CheckPwd(username string, pwd string) (bool, error) {
	user, err := u.UserRepository.FindUserByName(username)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.PassWord)
}
