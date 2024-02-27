package main

import (
	"bytes"
	"log"
	"golang.org/x/net/http2"
)

func main() {
    log.Println("Hello vulnerable")
   
    var buff bytes.Buffer
    var content = []byte("vulnerable")
    reader := bytes.NewReader(content)
    framer := http2.NewFramer(&buff, reader)

    log.Printf("%+v\n", framer)
}
