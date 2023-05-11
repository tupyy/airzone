package hvac

import (
	"bytes"
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

func GetData(host string, systemID, zoneID int) (Hvac, error) {
	p := Payload{SystemID: systemID, ZoneID: zoneID}

	payload, err := json.Marshal(p)
	if err != nil {
		return Hvac{}, err
	}

	url := fmt.Sprintf(base_url, host)
	resp, err := http.Post(url, jsonContentType, bytes.NewBuffer(payload))
	if err != nil {
		return Hvac{}, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Hvac{}, err
	}

	var hvac = Hvac{}
	if err := json.Unmarshal(data, &hvac); err != nil {
		return Hvac{}, err
	}

	return hvac, nil
}

func GetZoneNames(host string, systemID int) (map[string]int, error) {
	data, err := GetData(host, systemID, 0)
	if err != nil {
		return nil, err
	}
	names := make(map[string]int)
	for _, z := range data.Zones {
		names[strings.ToLower(z.Name)] = z.ID
	}
	return names, nil
}
