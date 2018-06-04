package main

import (
    _ "fmt"
    "github.com/ninjadotorg/handshake-services/fcm-service/config"
    "github.com/ninjadotorg/handshake-services/fcm-service/server"
)

func main() {
    config.Init()
    //db.Init()
    server.Init()
}
