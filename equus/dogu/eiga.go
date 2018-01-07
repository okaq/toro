// okaq web server
// equus dogu eiga
// bitmap sampler
// AQ <aq@okaq.com>
// 2018-01-08
package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    EIGA = "eiga.html"
    IMG = "../img"
)

func EigaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFIle(w,r,EIGA)
}

func main() {
    fmt.Printf("okaq equus eiga start\n%s\n", time.Now().String())
    http.HandleFunc("/", EigaHandler)
    http.ListenAndServe(":8080", nil)
}


