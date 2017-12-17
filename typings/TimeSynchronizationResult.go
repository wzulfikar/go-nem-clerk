// GO struct based on JSON file
// typings/json-sample/TimeSynchronizationResult.json
//
// This file was generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
package nemtypings

type TimeSynchronizationResult struct {
	Change            int64  `json:"change"`
	CurrentTimeOffset int64  `json:"currentTimeOffset"`
	DateTime          string `json:"dateTime"`
}
