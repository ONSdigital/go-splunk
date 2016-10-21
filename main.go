package main

import (
	"time"

	"github.com/ONSdigital/go-splunk/analytics"
	"github.com/ONSdigital/go-splunk/calendar"
)

func main() {
	calPer := 3600 //will be configurable
	cal := calendar.New()
	_ = analytics.New()

	calTicker := time.NewTicker(time.Second * time.Duration(calPer)).C

	//monitor channels to perform task when each ticker activates
	for {
		select {
		case <-calTicker:
			go cal.Check()
			//	case <-alyTicker:
			//		go aly.Check()
		}
	}

}
