// GO struct based on JSON file
// typings/json-sample/TransactionMetaData.json
//
// This file is generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
// → Sun Dec 17 17:11:03 MYT 2017
package nemtypings

type TransactionMetaData struct {
	Hash struct {
		Data string `json:"data"`
	} `json:"hash"`
	Height int64 `json:"height"`
	ID     int64 `json:"id"`
}
