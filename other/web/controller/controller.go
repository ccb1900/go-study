package controller

import log "github.com/sirupsen/logrus"

type Controller struct {
	log.Logger
}

func (c *Controller) Test() {
	c.Println("ccccc llll")
}
