package main

import (
	"bytes"
	"fmt"
	"github.com/algorand/go-algorand-sdk/client/algod"
	"github.com/algorand/go-algorand-sdk/client/kmd"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/future"
	"github.com/algorand/go-algorand-sdk/types"
)

//CHANGE ME
const kmdAddress = "http://localhost:7833"
const kmdToken = "b1105d6dc7192617a63acfc023d9a693aa5690dc20fbea40f571150bfc7d6339" //put here a key
const algodAddress = "http://localhost:8080"
const algodToken = "330b2e4fc9b20f4f89812cf87f1dabeb716d23e3f11aec97a61ff5f750563b78" //change this kmdAddress

func main() {
	fmt.Println("Hello Main Method")
}
