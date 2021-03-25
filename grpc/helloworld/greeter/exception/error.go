package exception

import log "github.com/sirupsen/logrus"

func Report(err error, s string) {
	if err != nil {
		log.Fatalf(s+":%v", err)
	}
}
