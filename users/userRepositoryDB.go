package users

import (
	"gomp/logger"

	"gorm.io/gorm"
)

type UserRepositoryDB struct {
	db *gorm.DB
}

type UserRepository interface {
	RegisterUser(Users) (Users, error)
	FindUser(string) (Users, error)
}

func NewUserRepositoryDB(client *gorm.DB) UserRepositoryDB {
	return UserRepositoryDB{client}
}

func (u UserRepositoryDB) RegisterUser(user Users) (Users, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		logger.Error("cannot register user " + err.Error())
		return user, err
	}
	return user, nil
}

func (u UserRepositoryDB) FindUser(username string) (Users, error) {
	err := u.db.Where("username=?", username).Find(&Users{}).Error
	if err != nil {
		logger.Error("cannot find user " + err.Error())
		return Users{}, err
	}
	return Users{}, nil
}
