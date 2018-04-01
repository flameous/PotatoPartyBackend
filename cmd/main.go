package main

import (
	"github.com/flameous/PotatoPartyBackend"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	s := potato.NewServer()
	s.Serve(nil)
}
