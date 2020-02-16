package viewmodel

import "os"

type StandLocator struct {
	Title  string
	Active string
	KeyEnv envConfig
}
type envConfig struct {
	ApiKey string
}

func NewStandLocator() StandLocator {

	result := StandLocator{
		Active: "standlocator",
		Title:  "Lemonade Stand Supply - Stand Locator",
		KeyEnv: envConfig{
			ApiKey: getEnv("MAPS_KEY", ""),
		},
	}
	return result
}

type StandCoordinate struct {
	Title     string  `json:"title"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
