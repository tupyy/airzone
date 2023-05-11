package hvac

type Hvac struct {
	Zones []Zone `json:"data"`
}

type Meta struct {
	ID   int    `json:"zoneID"`
	Name string `json:"name"`
}

type Parameters struct {
	On           int     `json:"on"`
	CoolSetPoint float64 `json:"coolsetpoint,omitempty"`
	CoolMaxTemp  float64 `json:"coolmaxtemp,omitemtpy"`
	CoolMinTemp  float64 `json:"coolmintemp,omitempty"`
	HeatSetPoint float64 `json:"heatsetpoint,omitempty"`
	HeatMaxTemp  float64 `json:"heatmaxtemp,omitempty"`
	HeatMinTemp  float64 `json:"heatmintemp,omitempty"`
}

type Zone struct {
	Meta
	Parameters
	Mode     int64   `json:"mode"`
	Modes    []int64 `json:"modes"`
	RoomTemp float64 `json:"roomTemp"`
	Humidity int64   `json:"humidity"`
}

type basePayload struct {
	SystemID int `json:"systemID"`
	ZoneID   int `json:"zoneID"`
}

type payload struct {
	basePayload
	Parameters
}
