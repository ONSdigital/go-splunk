package main

import (
	"flag"

	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-splunk/analytics"
	"github.com/ONSdigital/go-splunk/calendar"
)

func main() {
	log.Namespace = "go-splunk"

	calPer := flag.Int("cal-period", 300, "number of seconds to wait before checking for calendar updates")
	gaPer := flag.Int("ga-period", 10, "number of seconds to wait before checking for analytics updates")
	//check gaPer - if < 11 - show warning that this may use the majority of the users Quota
	// 4 is only if there is 1 request being made by this application, there are at least 2
	//`a single Google Analytics view (profile) has a daily quota limit of 10,000 requests per day.`
	flag.Parse()

	cal := calendar.New(calPer)
	ga := analytics.New(gaPer)

	//Wait for any ticker to send a message via its channel.
	//The channel will wait for its case to be selected, then become inactive
	//until it's ticker rolls over again.
	for {
		select {
		case <-cal.Ticker:
			go cal.Check()
		case <-ga.Ticker:
			go ga.Check()
		}
	}

}
