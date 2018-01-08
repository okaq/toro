// okaq web server
// equus dogu eiga
// bitmap sampler
// AQ <aq@okaq.com>
// 2018-01-08
package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
)

var (
    Images []string
)

const (
    EIGA = "eiga.html"
    IMG = "../img"
)

func EigaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,EIGA)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // api endpoint design
    // all app data can be served as json
    // list png image names in img dir
}

func Pop() {
    // populate image file names
    var err error
    f0, err := os.Open(IMG)
    if err != nil {
        fmt.Println(err)
    }
    Images, err := f0.Readdirnames(0)
    if err != nil {
        fmt.Println(err)
    }
    for index, name := range Images {
        fmt.Printf("File #%d:  %s\n", index, name)
    }
}

func main() {
    fmt.Printf("okaq equus eiga start\n%s\n", time.Now().String())
    Pop()
    http.HandleFunc("/", EigaHandler)
    http.HandleFunc("/list", ListHandler)
    http.ListenAndServe(":8080", nil)
}


