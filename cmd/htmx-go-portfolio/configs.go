package main

import (
	"encoding/json"
	"os"
)

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
