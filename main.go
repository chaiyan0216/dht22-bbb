package main

/*
#include "lib/bbb_dht_read.h"
#include "lib/bbb_dht_read.c"
#include "lib/bbb_mmio.c"
#include "lib/common_dht_read.c"
*/
import "C"
import (
    "flag"
    "fmt"
    "log"
    "net/http"
)

const (
    DHT22 = 22
    P9 = 1
    PIN12 = 28
)

var root = flag.String("root", "/home/cy/dht22-bbb/static", "Web Root Directory")
var port = flag.String("port", "8081", "Web Server Port")

func getHumAndTemp(w http.ResponseWriter, request *http.Request) {
    w.Header().Set("Content-Type", "text/json")

    hum := C.float(0.00)
    temp := C.float(0.00)
    C.bbb_dht_read(DHT22, P9, PIN12, &hum, &temp)

    fmt.Fprintf(w, "[%.2f, %.2f]", hum, temp)
}

func main() {
    flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*root)))
	http.Handle("/api/get", http.HandlerFunc(getHumAndTemp))

    err := http.ListenAndServe(":" + *port, nil)
	if err != nil {
		log.Panicln("ListenAndServe:", err)
	}
}
