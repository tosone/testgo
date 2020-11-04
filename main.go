package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("hello there")
		<-time.After(time.Second * 5)
	}
}
