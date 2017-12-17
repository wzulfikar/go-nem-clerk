// GO struct based on JSON file
// typings/json-sample/TransactionMetaData.json
package nemtypings

type TransactionMetaData struct {
	Hash   Hash  `json:"hash"`
	Height int64 `json:"height"`
	ID     int64 `json:"id"`
}
