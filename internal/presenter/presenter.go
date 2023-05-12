package presenter

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/tupyy/airzone/internal/hvac"
)

func Json(v interface{}) (string, error) {
	hvac, ok := v.(hvac.Hvac)
	if !ok {
		return "", errors.New("invalid model.want hvac.Hvac")
	}

	j, err := json.MarshalIndent(hvac.Zones, "", "\t")
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func Table(v interface{}) (string, error) {
	hvac, ok := v.(hvac.Hvac)
	if !ok {
		return "", errors.New("invalid model.want hvac.Hvac")
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"ZoneID", "Name", "On", "Mode", "CoolSetPoint °C", "HeatSetPoint °C", "Temperature °C", "Humidity %"})

	for _, z := range hvac.Zones {
		tw.AppendRow(table.Row{z.ID, z.Name, z.On, z.GetMode().String(), z.CoolSetPoint, z.HeatSetPoint, z.RoomTemp, z.Humidity})
	}

	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 3, Align: text.AlignRight, AlignHeader: text.AlignCenter, Transformer: func(val interface{}) string {
			v, _ := val.(int)
			switch v {
			case 0:
				return "off"
			case 1:
				return "on"
			default:
				return ""
			}
		}},
		{Number: 5, Align: text.AlignRight, AlignHeader: text.AlignCenter, Transformer: func(val interface{}) string { return fmt.Sprintf("%.2f", val) }},
		{Number: 6, Align: text.AlignRight, AlignHeader: text.AlignCenter, Transformer: func(val interface{}) string { return fmt.Sprintf("%.2f", val) }},
		{Number: 7, Align: text.AlignRight, AlignHeader: text.AlignCenter, Transformer: func(val interface{}) string { return fmt.Sprintf("%.2f", val) }},
	})
	return tw.Render(), nil
}
