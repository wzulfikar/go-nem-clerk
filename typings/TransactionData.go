package nemtypings

// TransactionData is `transaction` with its `meta` data
type TransactionData struct {
	Meta        TransactionMetaData `json:"meta"`
	Transaction Transaction         `json:"transaction"`
}
