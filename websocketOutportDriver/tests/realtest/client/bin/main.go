package main

import (
	"flag"
	"log"

	"github.com/ElrondNetwork/elrond-go-core/marshal"
	"github.com/ElrondNetwork/elrond-go-core/websocketOutportDriver/tests/realtest/client"
)

var (
	addr = flag.String("name", "client 0", "-")
	port = flag.Int("port", 21111, "-")
)

func main() {
	tc, err := client.NewTempClient(*addr, &marshal.JsonMarshalizer{})
	if err != nil {
		log.Fatal(err.Error())
	}

	tc.Run(*port)
}
