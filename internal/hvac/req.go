package hvac

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	base_url        = "http://%s/api/v1/hvac"
	jsonContentType = "application/json"
)

func GetData(ctx context.Context, host string, systemID, zoneID int) (Hvac, error) {
	p := basePayload{SystemID: systemID, ZoneID: zoneID}

	url := fmt.Sprintf(base_url, host)

	return do(ctx, http.MethodPost, url, p, func(data []byte) (Hvac, error) {
		var hvac = Hvac{}
		if err := json.Unmarshal(data, &hvac); err != nil {
			return Hvac{}, err
		}

		return hvac, nil
	})
}

func SetTemperature(ctx context.Context, host string, systemID, zoneID int, temperature float64) (Hvac, error) {
	mode, err := getMode(ctx, host, systemID)
	if err != nil {
		return Hvac{}, err
	}

	p := payload{
		basePayload: basePayload{
			SystemID: systemID,
			ZoneID:   zoneID,
		},
	}

	switch mode {
	case HeatingMode:
		p.HeatSetPoint = temperature
	case CoollingMode:
		p.CoolSetPoint = temperature
	}

	return set(ctx, host, p)
}

func SetMode(ctx context.Context, host string, systemID, zoneID int, mode Mode) (Hvac, error) {
	p := payload{
		basePayload: basePayload{
			SystemID: systemID,
			ZoneID:   zoneID,
		},
		Parameters: Parameters{
			Mode: int64(mode),
		},
	}

	return set(ctx, host, p)
}

func Start(ctx context.Context, host string, systemID, zoneID int, on bool) (Hvac, error) {
	action := 0
	if on {
		action = 1
	}
	p := payload{
		basePayload: basePayload{
			SystemID: systemID,
			ZoneID:   zoneID,
		},
		Parameters: Parameters{
			On: action,
		},
	}

	return set(ctx, host, p)
}

func GetZoneNames(ctx context.Context, host string, systemID int) (map[string]int, error) {
	data, err := GetData(ctx, host, systemID, 0)
	if err != nil {
		return nil, err
	}
	names := make(map[string]int)
	for _, z := range data.Zones {
		names[strings.ToLower(z.Name)] = z.ID
	}
	return names, nil
}

func getMode(ctx context.Context, host string, systemID int) (Mode, error) {
	data, err := GetData(ctx, host, systemID, 0)
	if err != nil {
		return 0, err
	}
	return Mode(data.Zones[0].Mode), nil
}

func set(ctx context.Context, host string, payload payload) (Hvac, error) {
	url := fmt.Sprintf(base_url, host)

	return do(ctx, http.MethodPut, url, payload, func(data []byte) (Hvac, error) {
		var hvac = Hvac{}
		if err := json.Unmarshal(data, &hvac); err != nil {
			return Hvac{}, err
		}

		return hvac, nil
	})
}

func do[T any](ctx context.Context, method, host string, payload interface{}, readFn func(data []byte) (T, error)) (T, error) {
	var emptyResponse T

	data, err := json.Marshal(payload)
	if err != nil {
		return emptyResponse, err
	}

	req, err := http.NewRequestWithContext(ctx, method, host, bytes.NewBuffer(data))
	if err != nil {
		return emptyResponse, err
	}
	req.Header.Add("ContentType", jsonContentType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return emptyResponse, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return emptyResponse, err
	}

	return readFn(content)
}
