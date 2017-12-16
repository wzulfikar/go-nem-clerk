# NEM Client (WIP)

- based on https://nemproject.github.io


```go
func main() {

    c, _ := nemclerk.NewClient("http://23.228.67.85:7890")

    // get all tx from given address
    tx, err := c.GetAllTransactions("NEM_ADDRESS_HERE", "", "")
    if err != nil {
        log.Fatal(err)
    }
    s, _ := prettyjson.Marshal(tx)
    fmt.Println(string(s))
}
```
