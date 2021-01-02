package main

import (
	"flag"
	"fmt"

	"github.com/Coderlane/go-minecraft-ping/client"
	"github.com/Coderlane/go-minecraft-ping/mcclient"
)

var server = flag.String("server", "localhost:25565",
	"The server to ping.")

func main() {
	fmt.Printf("Connecting to: %s\n", *server)
	rawcnt, err := client.NewClient(*server)
	if err != nil {
		fmt.Println(err)
		return
	}
	cnt, err := mcclient.NewMinecraftClient(rawcnt)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = cnt.Handshake(mcclient.ClientStateStatus)
	if err != nil {
		fmt.Println(err)
		return
	}
	status, err := cnt.Status()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Response: %s\n", status)
	err = cnt.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
