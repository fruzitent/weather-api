package entity

import "git.fruzit.pp.ua/weather/api/internal/shared/domain/value"

type User struct {
	Id   value.Id
	Mail value.Mail
}

func NewUser(id value.Id, mail value.Mail) *User {
	return &User{
		Id:   id,
		Mail: mail,
	}
}
