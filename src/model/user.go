package model

import (
	orm "text1/text1/src/db"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:password"`
}

var Users []User

//增
func (user User) Insert() (id int64, err error) {
	result := orm.Eloquent.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//删
func (user *User) Delete(id int64) (Reslut User, err error) {

	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}
	Reslut = *user
	return
}

//改
func (user *User) Update(id int64) (updateUser User, err error) {
	if err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

//查
func (user *User) Users() (users []User, err error) {
	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}
