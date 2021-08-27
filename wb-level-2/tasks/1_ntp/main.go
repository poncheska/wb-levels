package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

func main() {
	curTime := time.Now()
	exTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current time: %v\nExact time: %v\n", curTime, exTime)
}
