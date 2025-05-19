package entity

import (
	valueShared "git.fruzit.pp.ua/weather/api/internal/shared/domain/value"
	"git.fruzit.pp.ua/weather/api/pkg/weather/domain/value"
)

type Report struct {
	CreatedAt int64
	Id        valueShared.Id
	Location  value.Location
	Forecast  Forecast
}

func NewReport(
	createdAt int64,
	id valueShared.Id,
	location value.Location,
	forecast Forecast,
) Report {
	return Report{createdAt, id, location, forecast}
}
