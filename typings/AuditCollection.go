// GO struct based on JSON file
// typings/json-sample/AuditCollection.json
//
// This file was generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
package nemtypings

type AuditCollection struct {
	Most_recent []struct {
		Elapsed_time int64  `json:"elapsed-time"`
		Host         string `json:"host"`
		ID           int64  `json:"id"`
		Path         string `json:"path"`
		Start_time   int64  `json:"start-time"`
	} `json:"most-recent"`
	Outstanding []struct {
		Elapsed_time int64  `json:"elapsed-time"`
		Host         string `json:"host"`
		ID           int64  `json:"id"`
		Path         string `json:"path"`
		Start_time   int64  `json:"start-time"`
	} `json:"outstanding"`
}
