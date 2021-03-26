package exception

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func Report(err error, s string) {
	log.Errorf(s+" %v", err)
}

const dev = false

func Debug(a ...interface{}) {
	if dev {
		fmt.Println(a)
	}
}
