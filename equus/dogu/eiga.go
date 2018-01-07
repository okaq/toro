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

func ListHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // api endpoint design
    // all app data can be served as json
    // list png image names in img dir
}

func main() {
    fmt.Printf("okaq equus eiga start\n%s\n", time.Now().String())
    http.HandleFunc("/", EigaHandler)
    http.HandleFunc("/list", ListHandler)
    http.ListenAndServe(":8080", nil)
}


