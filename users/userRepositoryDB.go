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
	FindByID(id int) (Users, error)
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
	users := Users{}
	err := u.db.Where("username = ?", username).Find(&users).Error
	// fmt.Println("repo db", users)
	if err != nil {
		logger.Error("cannot find user " + err.Error())
		return users, err
	}
	return users, nil
}

func (u UserRepositoryDB) FindByID(id int) (Users, error) {
	var user Users
	var err error
	if err = u.db.Where("user_id = ?", id).Find(&user).Error; err != nil {
		logger.Error("Unexpected Error: " + err.Error())
		return user, err
	}
	return user, nil
}
