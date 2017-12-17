// GO struct based on JSON file
// typings/json-sample/UnconfirmedTransactionMetaDataPair.json
//
// This file was generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
package nemtypings

type UnconfirmedTransactionMetaDataPair struct {
	Data []struct {
		Meta        MetaData `json:"meta"`
		Transaction struct {
			Deadline   int64 `json:"deadline"`
			Fee        int64 `json:"fee"`
			OtherTrans struct {
				Amount   int64 `json:"amount"`
				Deadline int64 `json:"deadline"`
				Fee      int64 `json:"fee"`
				Message  struct {
					Payload string `json:"payload"`
					Type    int64  `json:"type"`
				} `json:"message"`
				Recipient string `json:"recipient"`
				Signer    string `json:"signer"`
				TimeStamp int64  `json:"timeStamp"`
				Type      int64  `json:"type"`
				Version   int64  `json:"version"`
			} `json:"otherTrans"`
			Signature  string        `json:"signature"`
			Signatures []interface{} `json:"signatures"`
			Signer     string        `json:"signer"`
			TimeStamp  int64         `json:"timeStamp"`
			Type       int64         `json:"type"`
			Version    int64         `json:"version"`
		} `json:"transaction"`
	} `json:"data"`
}
