package exception

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

const dev = true

func Report(err error, s string) {
	if dev {
		log.Errorf(s+" %v", err)
	}
}

func Debug(a ...interface{}) {
	if dev {
		fmt.Println(a)
	}
}
