package main

import (
    "fmt"
    "net/http"

    "./hrpwing"
)

func main() {
    p := &hrpwing.Proxy {
        Client: http.DefaultClient,
        BaseURL: "https://www.golang.org",
    }

    http.Handle("/", p)
    fmt.Println("Listening on port :3000")
    err := http.ListenAndServe(":3000", nil)
    panic(err)
}
