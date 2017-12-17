// GO struct based on JSON file
// typings/json-sample/TransactionMetaData.json
//
// This file was generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
package nemtypings

type TransactionMetaData struct {
	Hash   Hash  `json:"hash"`
	Height int64 `json:"height"`
	ID     int64 `json:"id"`
}
