package main

import (
	"flag"

	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-splunk/analytics"
	"github.com/ONSdigital/go-splunk/calendar"
)

func main() {
	log.Namespace = "go-splunk"

	calPer := flag.Int("cal-period", 3600, "number of seconds to wait before checking for calendar updates")
	flag.Parse()

	cal := calendar.New(calPer)
	_ = analytics.New()

	//Wait for any ticker to send a message via its channel.
	//The channel will wait for its case to be selected, then become inactive
	//until it's ticker rolls over again.
	for {
		select {
		case <-cal.Ticker:
			go cal.Check()
			//	case <-alyTicker:
			//		go aly.Check()
		}
	}

}
