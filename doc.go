// Go NEM Client watchs transactions from multisig account and signs it.
//
//
// ## package nemclient
//
// Sample code:
//	func main() {
//		c, _ := nemclient.NewClient("http://23.228.67.85:7890")
//		tx, err := c.GetAllTransactions("NEM_ADDRESS_HERE", "", "")
//		if err != nil {
//			log.Fatal(err)
//		}
//		s, _ := prettyjson.Marshal(tx)
//		fmt.Println(string(s))
//	}
//
// ## package nemrequests
//
// Wrapper for nemclient and nemparams.
//
// Method name convention `{HTTP Method}{endpoint name}`. e,g:
//   - `GetAllTransactions`
//   - `PostBlockOfHeight`
//
// ## package nemparams
//
// params  used in nemrequests. empty strings are omitted (using `nemrequests/client.go#omitEmptyParams`).
//
// ## package nemtypings
//
// typings for nem objects. files here are generated using ChimeraCoder/gojson via `typings/generate.sh`.
// You may need to adjust some of those generated files if it contains nested object.
// ie. `Block.go` has nested object []Transaction which can't be generated using generate.sh
//
// ## Environment Variables
//	- GO_NEM_CONFIG
//	- NEM_SERVER
//	- NEM_CLERK_ADDRESS
//
package nemclient
