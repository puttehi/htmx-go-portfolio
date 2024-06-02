package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// https://stackoverflow.com/a/48051946
type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalJSON unmarshals a JSON representation of SVGData from an alias (SVGCode/SVGItchIo/...) to its actual Go struct type
func (d *Duration) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return fmt.Errorf("Could not parse duration string %s: %s", value, err)
		}
		return nil
	default:
		return fmt.Errorf("Invalid duration %s. Want number (time.Duration value) or string (duration string like 10s)", value)
	}
}

type Config struct {
	PersonName            string   `json:"name"`
	NextNavbarAction      string   `json:"next_navbar_action"`
	ConfigRefreshInterval Duration `json:"config_refresh_interval"`
}

type WorkExperienceItem struct {
	RoleTitle string        `json:"roleTitle"`
	Company   string        `json:"company"`
	Date      string        `json:"date"`
	Details   []ItemDetails `json:"details"`
}

type LinkData struct {
	URL     string  `json:"url"`
	Text    string  `json:"text"`
	SVGData SVGData `json:"svg"` // Custom unmarshal from alias -> hard-coded data
}

type ProjectItem struct {
	Name     string        `json:"name"`
	Platform string        `json:"platform"`
	Links    []LinkData    `json:"links"`
	Video    string        `json:"video"`
	Details  []ItemDetails `json:"details"`
}

// ReadContentFromJSON reads a config file for content, unmarshaling to given type
//
// Example:
//
//	jsonAsStruct := MyUnmarshalableType{}
//	err = ReadContentFromJSON[MyUnmarshalableType]("my-config.json"), &jsonAsStruct)
//
//		if err != nil {
//	    slog.Default().Error("Could not read config file", "err", err)
//	    os.Exit(1)
//	}
func ReadContentFromJSON[T any](configFilePath string, unmarshalTo *T) error {
	b, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &unmarshalTo)
	if err != nil {
		return err
	}

	return nil
}
