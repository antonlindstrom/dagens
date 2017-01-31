package dagens

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// BaseURL is the base URL for the API.
const BaseURL = "http://api.dryg.net/dagar/v2.1"

// Response is the response structure returned from the API.
type Response struct {
	CacheTime string `json:"cachetid"`
	Version   string `json:"version"`
	Uri       string `json:"uri"`
	StartDate string `json:"startdatum"`
	StopDate  string `json:"stopdatum"`
	Days      []Day  `json:"dagar"`
}

// Day returns information about a day.
type Day struct {
	Date     string      `json:"datum"`
	Weekday  string      `json:"veckodag"`
	IsDayOff SwedishBool `json:"arbetsfri dag"`
	IsRedDay SwedishBool `json:"r\u00f6d dag"`
	Holiday  string      `json:"helgdag"`
	Names    []string    `json:"namnsdag"`
}

// SwedishBool is a mapping between Ja == true.
type SwedishBool string

// Bool returns true if s == Ja.
func (s SwedishBool) Bool() bool {
	return s == "Ja"
}

// Dagens is an interface for fetching the date.
type Dagens interface {
	Date(time.Time) (*Response, error)
}

// Date fetches the information about today from upstream API.
func Date(t time.Time) (*Response, error) {
	url := fmt.Sprintf("%s/%d/%02d/%02d", BaseURL, t.Year(), t.Month(), t.Day())
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data *Response
	err = json.Unmarshal(body, &data)

	return data, err
}
