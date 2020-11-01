package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"shut-me-down/pkg/auth"
	"shut-me-down/pkg/fetch"
	// "shut-me-down/pkg/fetch"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	/*
		https://www.honeybadger.io/blog/golang-logging/
		examples
		log.Debug("Useful debugging information.")
		log.Info("Something noteworthy happened!")
		log.Warn("You should probably take a look at this.")
		log.Error("Something failed but I'm not quitting.")
	*/
}

func main() {
	// var flagvar = flag.Int("flagname", 1234, "help message for flagname")
	// var dryrun = flag.Bool("dryrun", false, "do a test run to see which instances would be shutdown")

	flag.Parse()

	session := auth.Login()

	// result := fetch.GetSmdInstances(session)

	fetch.GetSmdInstances(session)

	// fmt.Println(result)
	//fetch.Get_ec2_tags(session)
}
