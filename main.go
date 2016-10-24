package main

import (
	"time"

	"github.com/ONSdigital/go-splunk/analytics"
	"github.com/ONSdigital/go-splunk/calendar"
)

func main() {
	calPer := 10 //will be configurable
	cal := calendar.New()
	_ = analytics.New()

	calTicker := time.NewTicker(time.Second * time.Duration(calPer)).C

	//Wait for any ticker to send a message via its channel.
	//The channel will wait for its case to be selected, then become inactive
	//until it's ticker rolls over again.
	for {
		select {
		case <-calTicker:
			go cal.Check()
			//	case <-alyTicker:
			//		go aly.Check()
		}
	}

}
