// GO struct based on JSON file
// typings/json-sample/Transaction.json
//
// This file is generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
// → Sun Dec 17 17:11:03 MYT 2017
package nemtypings

type Transaction struct {
	Deadline      int64  `json:"deadline"`
	Fee           int64  `json:"fee"`
	Mode          int64  `json:"mode"`
	RemoteAccount string `json:"remoteAccount"`
	Signature     string `json:"signature"`
	Signer        string `json:"signer"`
	TimeStamp     int64  `json:"timeStamp"`
	Type          int64  `json:"type"`
	Version       int64  `json:"version"`
}
