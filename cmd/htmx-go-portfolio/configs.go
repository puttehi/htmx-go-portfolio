package main

import (
	"encoding/json"
	"os"
)

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
