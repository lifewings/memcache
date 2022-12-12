package main

import (
	"log"

	"memcache/cmd"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile | log.LUTC)
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
