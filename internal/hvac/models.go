package hvac

type Hvac struct {
	Zones []Zone `json:"data"`
}

type Zone struct {
	Name         string  `json:"name"`
	On           int     `json:"on"`
	CoolSetPoint float64 `json:"coolsetpoint"`
	CoolMaxTemp  float64 `json:"coolmaxtemp"`
	CoolMinTemp  float64 `json:"coolmintemp"`
	HeatSetPoint float64 `json:"heatsetpoint"`
	HeatMaxTemp  float64 `json:"heatmaxtemp"`
	HeatMinTemp  float64 `json:"heatmintemp"`
	Mode         int64   `json:"mode"`
	Modes        []int64 `json:"modes"`
	RoomTemp     float64 `json:"roomTemp"`
	Humidity     int64   `json:"humidity"`
}

type Payload struct {
	SystemID int `json:"systemID"`
	ZoneID   int `json:"zoneID"`
}
