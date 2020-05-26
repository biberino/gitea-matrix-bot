package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	conf := LoadConfiguration("conf.json")
	fmt.Println("Started.")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(conf.Address, nil))
}
