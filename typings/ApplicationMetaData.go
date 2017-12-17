// GO struct based on JSON file
// typings/json-sample/ApplicationMetaData.json
//
// This file is generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
// → Sun Dec 17 17:11:03 MYT 2017
package nemtypings

type ApplicationMetaData struct {
	Application string `json:"application"`
	CurrentTime int64  `json:"currentTime"`
	Signer      string `json:"signer"`
	StartTime   int64  `json:"startTime"`
	Version     string `json:"version"`
}
