// GO struct based on JSON file
// typings/json-sample/RequestPrepareAnnounce.json
//
// This file is generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
// → Sun Dec 17 17:11:03 MYT 2017
package nemtypings

type RequestPrepareAnnounce struct {
	PrivateKey  string `json:"privateKey"`
	Transaction struct {
		Deadline     int64  `json:"deadline"`
		Fee          int64  `json:"fee"`
		OtherAccount string `json:"otherAccount"`
		OtherHash    struct {
			Data string `json:"data"`
		} `json:"otherHash"`
		Signer    string `json:"signer"`
		TimeStamp int64  `json:"timeStamp"`
		Type      int64  `json:"type"`
		Version   int64  `json:"version"`
	} `json:"transaction"`
}
