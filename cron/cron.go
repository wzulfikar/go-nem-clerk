package main

import (
	"fmt"
	"log"
	"os"

	nemclerk "github.com/wzulfikar/go-nem-client"
	nemrequests "github.com/wzulfikar/go-nem-client/requests"

	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"

	prettyjson "github.com/hokaccha/go-prettyjson"
)

func main() {
	err := godotenv.Load(os.Getenv("GO_NEM_CONFIG"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c, err := nemclerk.DefaultClient()
	if err != nil {
		panic(err)
	}

	s := gocron.NewScheduler()

	clerkAddr := os.Getenv("NEM_CLERK_ADDR")
	fmt.Println("Watching unsigned tx at address " + clerkAddr)
	s.Every(3).Seconds().Do(func() { signUnsignedTx(c, clerkAddr) })

	<-s.Start()
}

func signUnsignedTx(c *nemrequests.Client, clerkAddr string) {
	tx, err := c.GetUnconfirmedTransactions(clerkAddr)
	if err != nil {
		log.Fatal(err)
	}

	if len(tx.Data) > 0 {
		fmt.Println("Unconfirmed tx: ", len(tx.Data))

		s, _ := prettyjson.Marshal(tx)
		fmt.Println(string(s))
	} else {
		fmt.Println("Waiting for unsigned tx..")
	}
}
