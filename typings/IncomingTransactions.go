// GO struct based on JSON file
// typings/json-sample/IncomingTransactions.json
//
// This file is generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
// → Sun Dec 17 17:11:03 MYT 2017
package nemtypings

type IncomingTransactions struct {
	Data []struct {
		Meta struct {
			Hash struct {
				Data string `json:"data"`
			} `json:"hash"`
			Height int64 `json:"height"`
			ID     int64 `json:"id"`
		} `json:"meta"`
		Transaction struct {
			Amount   int64 `json:"amount"`
			Deadline int64 `json:"deadline"`
			Fee      int64 `json:"fee"`
			Message  struct {
				Payload string `json:"payload"`
				Type    int64  `json:"type"`
			} `json:"message"`
			Recipient string `json:"recipient"`
			Signature string `json:"signature"`
			Signer    string `json:"signer"`
			TimeStamp int64  `json:"timeStamp"`
			Type      int64  `json:"type"`
			Version   int64  `json:"version"`
		} `json:"transaction"`
	} `json:"data"`
}
