package service

type IProbe interface {
	IsHealthy() error
	IsReady() error
}
