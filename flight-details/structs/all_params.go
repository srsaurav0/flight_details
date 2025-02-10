package structs

// FlightSearchParams defines the search filters
type FlightSearchParams struct {
	FlightNum          string  `json:"FlightNum,omitempty"`
	DestCountry        string  `json:"DestCountry,omitempty"`
	OriginWeather      string  `json:"OriginWeather,omitempty"`
	OriginCityName     string  `json:"OriginCityName,omitempty"`
	AvgTicketPrice     float64 `json:"AvgTicketPrice,omitempty"`
	DistanceMiles      float64 `json:"DistanceMiles,omitempty"`
	FlightDelay        bool    `json:"FlightDelay,omitempty"`
	DestWeather        string  `json:"DestWeather,omitempty"`
	Dest               string  `json:"Dest,omitempty"`
	FlightDelayType    string  `json:"FlightDelayType,omitempty"`
	OriginCountry      string  `json:"OriginCountry,omitempty"`
	DayOfWeek          int     `json:"dayOfWeek,omitempty"`
	DistanceKilometers float64 `json:"DistanceKilometers,omitempty"`
	TravelTime         string  `json:"timestamp"` // Mandatory Field
	DestLocationLat    string  `json:"DestLocationLat,omitempty"`
	DestLocationLon    string  `json:"DestLocationLon,omitempty"`
	DestAirportID      string  `json:"DestAirportID,omitempty"`
	Carrier            string  `json:"Carrier,omitempty"`
	Cancelled          bool    `json:"Cancelled,omitempty"`
	FlightTimeMin      float64 `json:"FlightTimeMin,omitempty"`
	Origin             string  `json:"Origin,omitempty"`
	OriginLocationLat  string  `json:"OriginLocationLat,omitempty"`
	OriginLocationLon  string  `json:"OriginLocationLon,omitempty"`
	DestRegion         string  `json:"DestRegion,omitempty"`
	OriginAirportID    string  `json:"OriginAirportID,omitempty"`
	OriginRegion       string  `json:"OriginRegion,omitempty"`
	DestCityName       string  `json:"DestCityName,omitempty"`
	FlightTimeHour     float64 `json:"FlightTimeHour,omitempty"`
	FlightDelayMin     int     `json:"FlightDelayMin,omitempty"`
}
