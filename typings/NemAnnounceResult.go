// This file is auto-generated using go-nem-clerk/typings/generate.sh\n
package nemtypings

type NemAnnounceResult struct {
	Code                 int64 `json:"code"`
	InnerTransactionHash struct {
		Data string `json:"data"`
	} `json:"innerTransactionHash"`
	Message         string `json:"message"`
	TransactionHash struct {
		Data string `json:"data"`
	} `json:"transactionHash"`
	Type int64 `json:"type"`
}
