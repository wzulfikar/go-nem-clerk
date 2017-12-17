// GO struct based on JSON file
// typings/json-sample/ApplicationMetaData.json
//
// This file was generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
package nemtypings

type ApplicationMetaData struct {
	Application string `json:"application"`
	CurrentTime int64  `json:"currentTime"`
	Signer      string `json:"signer"`
	StartTime   int64  `json:"startTime"`
	Version     string `json:"version"`
}
