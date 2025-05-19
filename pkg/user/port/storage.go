package port

import "git.fruzit.pp.ua/weather/api/pkg/user/domain/entity"

type Storage interface {
	SaveUser(user entity.User) error
}
