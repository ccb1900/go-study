package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.Println(http.ListenAndServe(":1267", nil))
}
