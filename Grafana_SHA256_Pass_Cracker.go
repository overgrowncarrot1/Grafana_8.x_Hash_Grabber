package main

import (
    b64 "encoding/base64"
    hex "encoding/hex"
    "fmt"
    "log"
    "os"
    
)

func main() {
    var password,salt string
    fmt.Print("Hash: \n")
    fmt.Scanln(&password)
    fmt.Print("Salt: \n")
    fmt.Scanln(&salt)

    decoded_hash, _ := hex.DecodeString(password)
    hash64 := b64.StdEncoding.EncodeToString([]byte(decoded_hash))
    salt64 := b64.StdEncoding.EncodeToString([]byte(salt))
    
    f, err := os.Create("hash.txt")
    if err != nil {
    log.Fatal(err)
    }
        defer f.Close()
        _, err2 := f.WriteString("sha256:10000:" + salt64 + ":" + hash64 + "\n")
        if err2 != nil {
        log.Fatal(err2)
    }
        fmt.Println("done")
}
