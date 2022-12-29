package repository

import (
	"gaunRumaRestApi/config/db"
	"gaunRumaRestApi/entity"
)

type UserRepository interface {
	GetAllUsers() (*[]entity.User, error)
	GetUserByUserName(userName string) (*entity.User, error)
	CreateUser(userName string, password string) (*entity.User, error)
	UpdateUser(id int64, userName string, password string) (*entity.User, error)
	DeleteUser(id int) (bool, error)
}

type UserRepositoryImpl struct {
	dbHandler *db.Handler
}

// CreateUser implements UserRepository
func (u *UserRepositoryImpl) CreateUser(userName string, password string) (*entity.User, error) {
	user := entity.User{UserName: userName, UserPassword: password}
	result := u.dbHandler.DB.Create(&user)

	if result.Error != nil {
		return &entity.User{}, result.Error
	}

	return &user, nil
}

// DeletUser implements UserRepository
func (u *UserRepositoryImpl) DeleteUser(id int) (bool, error) {
	var user entity.User
	resp := u.dbHandler.DB.Delete(&user, id)
	if resp.Error != nil {
		return false, resp.Error
	}

	return true, nil
}

// GetAllUsers implements UserRepository
func (u *UserRepositoryImpl) GetAllUsers() (*[]entity.User, error) {
	var arrUser []entity.User
	var user entity.User
	resp := u.dbHandler.DB.Find(&user)

	if resp.Error != nil {
		return &arrUser, resp.Error
	}
	rows, err := resp.Rows()
	if err != nil {
		return &arrUser, err

	}
	for rows.Next() {
		err = rows.Scan(&user)
		if err != nil {
			return &arrUser, err
		}
		arrUser = append(arrUser, user)
	}
	return &arrUser, nil
}

// GetUserByUserName implements UserRepository
func (u *UserRepositoryImpl) GetUserByUserName(userName string) (*entity.User, error) {
	var user entity.User
	err := u.dbHandler.DB.Where("user_name = ?", userName).First(&user).Error
	return &user, err
}

// UpdateUser implements UserRepository
func (*UserRepositoryImpl) UpdateUser(id int64, userName string, password string) (*entity.User, error) {
	panic("unimplemented")
}

func NewUserRepository(dbHandler *db.Handler) UserRepository {
	return &UserRepositoryImpl{
		dbHandler: dbHandler,
	}
}
