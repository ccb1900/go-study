package exception

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

const dev = false

func Report(err error, s string) {
	if dev {
		log.Fatalf(s+" %v", err)
	}
}

func Debug(a ...interface{}) {
	if dev {
		fmt.Println(a)
	}
}
