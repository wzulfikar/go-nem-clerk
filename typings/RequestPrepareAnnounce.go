// GO struct based on JSON file
// typings/json-sample/RequestPrepareAnnounce.json
//
// This file was generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
package nemtypings

type RequestPrepareAnnounce struct {
	PrivateKey  string `json:"privateKey"`
	Transaction struct {
		Deadline     int64  `json:"deadline"`
		Fee          int64  `json:"fee"`
		OtherAccount string `json:"otherAccount"`
		OtherHash    Hash   `json:"otherHash"`
		Signer       string `json:"signer"`
		TimeStamp    int64  `json:"timeStamp"`
		Type         int64  `json:"type"`
		Version      int64  `json:"version"`
	} `json:"transaction"`
}
