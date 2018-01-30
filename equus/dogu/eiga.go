// okaq web server
// equus dogu eiga
// bitmap sampler
// AQ <aq@okaq.com>
// 2018-01-08
package main

import (
    "bufio"
    "fmt"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
)

var (
    Images []string
    Index int
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
    Images, err = f0.Readdirnames(0)
    if err != nil {
        fmt.Println(err)
    }
    r0 := bufio.NewReader(os.Stdin)
    for index, name := range Images {
        fmt.Printf("Select the file number to sample\n")
        fmt.Printf("\n(%d):  %s\n", index, name)
        fmt.Printf("\nPress input key:\n> ")
    }
    s0, _ := r0.ReadString('\n')
    // fmt.Printf(s0)
    s1 := strings.TrimSuffix(s0, "\n")
    // fmt.Println(s1)
    i0, _ := strconv.Atoi(s1)
    // fmt.Println(i0, Images[i0])
    Index = i0
    fmt.Printf("Input key pressed is: %d. File to be sampled: %s\n", i0, Images[i0])
    fmt.Println("File input saved. Ready to begin web serving...")
}

func SelectHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // write image file name
    // as plain text string
    w.Header().Set("Content-type", "text/plain")
    w.Write([]byte(Images[Index]))
    // logging request interface
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // http.StripPrefix("/img/", http.FileServer(http.Dir(IMG)))
    fmt.Println(r.URL.Path)
    s0 := strings.Split(r.URL.Path, "/")
    fmt.Println(s0)
    s1 := fmt.Sprintf("%s/%s", IMG, s0[2])
    fmt.Println(s1)
    http.ServeFile(w,r,s1)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func ProcHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func main() {
    fmt.Printf("okaq equus eiga start\n%s\n", time.Now().String())
    Pop()
    http.HandleFunc("/", EigaHandler)
    http.HandleFunc("/list", ListHandler)
    http.HandleFunc("/a", SelectHandler)
    http.HandleFunc("/img/", ImageHandler)
    http.HandleFunc("/save", SaveHandler)
    http.HandleFunc("/proc", ProcHandler)
    http.ListenAndServe(":8080", nil)
}


