// GO struct based on JSON file
// typings/json-sample/Error.json
//
// This file is generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
// → Sun Dec 17 17:11:03 MYT 2017
package nemtypings

type Error struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	Status    int64  `json:"status"`
	TimeStamp int64  `json:"timeStamp"`
}
