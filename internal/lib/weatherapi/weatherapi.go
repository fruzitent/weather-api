package weatherapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const BASE_URL = "https://api.weatherapi.com/v1"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Message)
}

type Config struct {
	Secret string
}

type Weatherapi struct {
	client *http.Client
	Secret string
}

func NewWeatherapi(secret string) (*Weatherapi, error) {
	return &Weatherapi{
		client: &http.Client{},
		Secret: secret,
	}, nil
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzId           string  `json:"tz_id"`
	LocaltimeEpoch int     `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type Current struct {
	LastUpdatedEpoch int       `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TempC            float64   `json:"temp_c"`
	TempF            float64   `json:"temp_f"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMph          float64   `json:"wind_mph"`
	WindKph          float64   `json:"wind_kph"`
	WindDegree       int       `json:"wind_degree"`
	WindDir          string    `json:"wind_dir"`
	PressureMb       float64   `json:"pressure_mb"`
	PressureIn       float64   `json:"pressure_in"`
	PrecipMm         float64   `json:"precip_mm"`
	PrecipIn         float64   `json:"precip_in"`
	Humidity         int       `json:"humidity"`
	Cloud            int       `json:"cloud"`
	FeelslikeC       float64   `json:"feelslike_c"`
	FeelslikeF       float64   `json:"feelslike_f"`
	WindchillC       float64   `json:"windchill_c"`
	WindchillF       float64   `json:"windchill_f"`
	HeatindexC       float64   `json:"heatindex_c"`
	HeatindexF       float64   `json:"heatindex_f"`
	DewpointC        float64   `json:"dewpoint_c"`
	DewpointF        float64   `json:"dewpoint_f"`
	VisKm            float64   `json:"vis_km"`
	VisMiles         float64   `json:"vis_miles"`
	Uv               float64   `json:"uv"`
	GustMph          float64   `json:"gust_mph"`
	GustKph          float64   `json:"gust_kph"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type RealtimeReq struct {
	Q    string `json:"q"`
	Lang string `json:"lang"`
}

type RealtimeRes struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

func (a *Weatherapi) Realtime(command *RealtimeReq) (*RealtimeRes, error) {
	uri, err := url.JoinPath(BASE_URL, "current.json")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	values.Add("key", a.Secret)
	values.Add("q", command.Q)
	req.URL.RawQuery = values.Encode()

	res, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		commandRes := &RealtimeRes{}
		if err := json.NewDecoder(res.Body).Decode(commandRes); err != nil {
			return nil, err
		}
		return commandRes, nil
	case http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden:
		errRes := &Error{}
		if err := json.NewDecoder(res.Body).Decode(errRes); err != nil {
			return nil, err
		}
		return nil, errRes
	}

	return nil, &Error{
		Code:    res.StatusCode,
		Message: res.Status,
	}
}
