// GO struct based on JSON file
// typings/json-sample/MosaicDefinitionCreationTransaction.json
//
// This file is generated using gojson via go-nem-client/typings/generate.sh
// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING
// → Sun Dec 17 17:11:03 MYT 2017
package nemtypings

type MosaicDefinitionCreationTransaction struct {
	CreationFee      int64  `json:"creationFee"`
	CreationFeeSink  string `json:"creationFeeSink"`
	Deadline         int64  `json:"deadline"`
	Fee              int64  `json:"fee"`
	MosaicDefinition struct {
		Creator     string `json:"creator"`
		Description string `json:"description"`
		ID          struct {
			Name        string `json:"name"`
			NamespaceID string `json:"namespaceId"`
		} `json:"id"`
		Levy struct {
			Fee      int64 `json:"fee"`
			MosaicID struct {
				Name        string `json:"name"`
				NamespaceID string `json:"namespaceId"`
			} `json:"mosaicId"`
			Recipient string `json:"recipient"`
			Type      int64  `json:"type"`
		} `json:"levy"`
		Properties []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"properties"`
	} `json:"mosaicDefinition"`
	Signature string `json:"signature"`
	Signer    string `json:"signer"`
	TimeStamp int64  `json:"timeStamp"`
	Type      int64  `json:"type"`
	Version   int64  `json:"version"`
}
