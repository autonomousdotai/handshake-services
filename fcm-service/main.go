package main

import (
    _ "fmt"
    "github.com/ninjadotorg/handshake-services/ipfs-service/config"
    "github.com/ninjadotorg/handshake-services/ipfs-service/server"
)

func main() {
    config.Init()
    //db.Init()
    server.Init()
}
