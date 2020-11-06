package models

type User struct {
	NickName string
	Email string
	Password string `xorm:"varchar(50)"`
}
