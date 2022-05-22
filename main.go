package main

import (
	"crypto/tls"
	"fmt"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.New()
	log := logrus.New()
	conn, err := tls.Dial("tcp", "901cf052-ee43-4085-950d-a58014f2051e-ls.logit.io:23191", &tls.Config{RootCAs: nil})
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "ti-3-law"}))
	log.Hooks.Add(hook)
	ctx := log.WithFields(logrus.Fields{
		"method": "main",
	})

	r.GET("/log", func(c *gin.Context) {
		messageQuery := c.DefaultQuery("message", "")
		ctx.Info(fmt.Sprintf("Get request with parameter %s", messageQuery))

		c.Data(200, "application/json; charset=utf-8", []byte(messageQuery))
	})
	r.Run()
}
