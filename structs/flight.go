package models

type Flight struct {
	FlightNum       string  `json:"FlightNum"`
	Carrier         string  `json:"Carrier"`
	DestCityName    string  `json:"DestCityName"`
	Timestamp       string  `json:"timestamp"`
	AvgTicketPrice  float64 `json:"AvgTicketPrice"`
	DestCountry     string  `json:"DestCountry"`
	OriginWeather   string  `json:"OriginWeather"`
	OriginCityName  string  `json:"OriginCityName"`
	DistanceMiles   float64 `json:"DistanceMiles"`
	FlightDelay     bool    `json:"FlightDelay"`
	DestWeather     string  `json:"DestWeather"`
	Dest            string  `json:"Dest"`
	FlightDelayType string  `json:"FlightDelayType"`
	OriginCountry   string  `json:"OriginCountry"`
	DayOfWeek       int     `json:"dayOfWeek"`
	DistanceKm      float64 `json:"DistanceKilometers"`
	DestAirportID   string  `json:"DestAirportID"`
	Cancelled       bool    `json:"Cancelled"`
	FlightTimeMin   float64 `json:"FlightTimeMin"`
	Origin          string  `json:"Origin"`
	DestRegion      string  `json:"DestRegion"`
	OriginAirportID string  `json:"OriginAirportID"`
	OriginRegion    string  `json:"OriginRegion"`
	FlightTimeHour  float64 `json:"FlightTimeHour"`
	FlightDelayMin  int     `json:"FlightDelayMin"`

	// Nested JSON objects for locations
	DestLocation struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	} `json:"DestLocation"`

	OriginLocation struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	} `json:"OriginLocation"`
}
