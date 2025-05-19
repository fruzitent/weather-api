package entity

import (
	valueShared "git.fruzit.pp.ua/weather/api/internal/shared/domain/value"
	"git.fruzit.pp.ua/weather/api/pkg/weather/domain/value"
)

type Report struct {
	Id          valueShared.Id
	Location    value.Location
	Measurement value.Measurement
}

func NewReport(id valueShared.Id, location value.Location, measurement value.Measurement) (*Report, error) {
	return &Report{id, location, measurement}, nil
}
