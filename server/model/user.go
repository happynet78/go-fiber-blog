package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name" gorm:"not null comment:first_name size:80"`
	LatName   string `json:"lat_name" gorm:"not null comment:lat_name size:80"`
	Email     string `json:"email" gorm:"not null comment:email size:255"`
	Password  []byte `json:"-"`
	Phone     string `json:"phone"`
}

func (user *User) SetPassword(password string) {
	bashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = bashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
