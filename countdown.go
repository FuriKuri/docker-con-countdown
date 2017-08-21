package main

import (
	"flag"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"time"
)

func main() {
	live := flag.Bool("live", false, "live modus")
	flag.Parse()
	dockerCon2017 := time.Date(2017, 10, 16, 7, 0, 0, 0, time.UTC)
	if *live {
		gocron.Every(1).Second().Do(func() { printRemainingTime(dockerCon2017) })
		<-gocron.Start()
	} else {
		printRemainingTime(dockerCon2017)
	}
}

func printRemainingTime(t time.Time) {
	secondsToDockerCon := int64(time.Until(t).Seconds())
	days := secondsToDockerCon / (24 * 60 * 60)
	secondsToDockerCon = secondsToDockerCon - 24*60*60*days
	hours := secondsToDockerCon / (60 * 60)
	secondsToDockerCon = secondsToDockerCon - 60*60*hours
	minutes := secondsToDockerCon / 60
	secondsToDockerCon = secondsToDockerCon - 60*minutes
	seconds := secondsToDockerCon

	fmt.Printf("\rTime until DockerCon Europe 2017: %d days %d hours %d minutes %d seconds",
		days, hours, minutes, seconds)
}
