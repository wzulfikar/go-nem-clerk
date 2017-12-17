// GO struct based on JSON file
// typings/json-sample/Block.json
//
// This file was generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
package nemtypings

type Block struct {
	Height        int64         `json:"height"`
	PrevBlockHash Hash          `json:"prevBlockHash"`
	Signature     string        `json:"signature"`
	Signer        string        `json:"signer"`
	TimeStamp     int64         `json:"timeStamp"`
	Transactions  []Transaction `json:"transactions"`
	Type          int64         `json:"type"`
	Version       int64         `json:"version"`
}
