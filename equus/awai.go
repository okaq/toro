// okaq web server
// horse swan
// AQ <aq@okaq.com>
// 2018-01-06
package main

import (
    "fmt"
    "net/http"
    "time"
)

const (
    INDEX = "zwai.html"
)

func AwaiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w,r,INDEX)
}

func main() {
    fmt.Printf("okaq equus started at\n%s\n", time.Now().String())
    http.HandleFunc("/", AwaiHandler)
    http.ListenAndServe(":8080", nil)
}


