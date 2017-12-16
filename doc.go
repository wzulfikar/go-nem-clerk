// Go NEM Clerk watchs transactions from multisig account and signs it.
//
// Sample code:
//	func main() {
//		c, _ := nemclerk.NewClient("http://23.228.67.85:7890")
//		tx, err := c.GetAllTransactions("NEM_ADDRESS_HERE", "", "")
//		if err != nil {
//			log.Fatal(err)
//		}
//		s, _ := prettyjson.Marshal(tx)
//		fmt.Println(string(s))
//	}
package nemclerk
