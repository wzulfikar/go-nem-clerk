// This file is auto-generated using go-nem-clerk/typings/generate.sh\n
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
