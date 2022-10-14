package main

import (
	"fmt"
	"foundation"
	"net/http"
)

func main() {
	fmt.Println("Canteros Web v0.0.1 🏗️")
	handler := foundation.NewHandler()
	err := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/canteros.digital/fullchain.pem", "/etc/letsencrypt/live/canteros.digital/privkey.pem", handler)
	if err != nil {
		err = http.ListenAndServe(":80", handler)
		if err != nil {
			fmt.Printf("%s", err.Error())
		}
	}
}
