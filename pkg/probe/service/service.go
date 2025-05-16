package service

type IService interface {
	IsHealthy() error
	IsReady() error
}
