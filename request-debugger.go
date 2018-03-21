package main

import (
    "fmt"
    "net/http"
    "sort"
    "bytes"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    var requestHeaders []string
    for k := range r.Header {
        requestHeaders = append(requestHeaders, k)
    }
    sort.Strings(requestHeaders)

    writer := bytes.Buffer{}
    fmt.Fprintln(&writer, "============ NEW REQUEST ============")

    fmt.Fprintln(&writer, "request.RequestURI:", r.RequestURI)
    fmt.Fprintln(&writer, "request.Host:", r.Host)
    fmt.Fprintln(&writer, "request.Method:", r.Method)
    fmt.Fprintln(&writer, "request.RemoteAddr:", r.RemoteAddr)
    fmt.Fprintln(&writer, "request.TLS:", r.TLS)

    fmt.Fprintln(&writer, "Request Headers:")
    for _, k := range requestHeaders {
        fmt.Fprintln(&writer, "\t", k, ":", r.Header[k])
    }

    fmt.Println(writer.String())

    w.Write([]byte("Request debugger (" + (time.Now()).Format(time.RFC3339) + ")"))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}