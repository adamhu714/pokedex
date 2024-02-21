package main

import (
	"net/http"
)

func commandMapF() error {
	http.Get("https://google.com")
	return nil
}

func commandMapB() error {
	return nil
}
