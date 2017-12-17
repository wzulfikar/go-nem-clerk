// GO struct based on JSON file
// typings/json-sample/NemAnnounceResult.json
//
// This file is generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
// → Sun Dec 17 17:11:03 MYT 2017
package nemtypings

type NemAnnounceResult struct {
	Code                 int64  `json:"code"`
	InnerTransactionHash Hash   `json:"innerTransactionHash"`
	Message              string `json:"message"`
	TransactionHash      Hash   `json:"transactionHash"`
	Type                 int64  `json:"type"`
}
