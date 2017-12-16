package nemrequests

import (
	"log"
	"os"
)

func logger(args ...interface{}) {
	if os.Getenv("ENABLE_LOG") == "true" {
		log.Println(args)
	}
}
