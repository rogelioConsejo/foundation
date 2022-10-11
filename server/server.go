package main

import (
	"fmt"
	"foundation"
	"net/http"
)

func main() {
	handler := foundation.NewHandler()
	err := http.ListenAndServe("localhost:1234", handler)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
}
