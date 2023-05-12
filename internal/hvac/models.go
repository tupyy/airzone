package hvac

type Mode int

const (
	StopMode Mode = 1 + iota
	CoollingMode
	HeatingMode
	VentilationMode
	Dehumidification
)

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
	Mode         int64   `json:"mode"`
}

type Zone struct {
	Meta
	Parameters
	Modes    []int64 `json:"modes,omitempty"`
	RoomTemp float64 `json:"roomTemp"`
	Humidity int64   `json:"humidity"`
}

func (z Zone) GetMode() Mode {
	return Mode(z.Mode)
}

type basePayload struct {
	SystemID int `json:"systemID"`
	ZoneID   int `json:"zoneID"`
}

type payload struct {
	basePayload
	Parameters
}

func (p payload) SetMode(m Mode) {
	p.Mode = int64(m)
}
