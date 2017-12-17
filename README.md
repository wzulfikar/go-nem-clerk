# NEM Client (WIP)

- based on https://nemproject.github.io



### 1. nemclient
- used to create instance of `nemrequests.Client`

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

### 2. nemrequests
- wrapper for path & params

### 3. nemparams
- request params

### 4. nemtypings
- typings (response, etc)