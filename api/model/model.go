package model

type GeoData struct {
	IP string `json:"ip"`
	Location string `json:"city"`
}

type TempResponse struct {
	CurrentRes Current `json:"current"`
}

type Current struct {
	TempCelsius float64 `json:"temp_c"`
}
