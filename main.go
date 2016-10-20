package main

import (
	"github.com/ONSdigital/go-splunk/analytics"
	"github.com/ONSdigital/go-splunk/calendar"
)

func main() {
	cal := calendar.New()
	_ = analytics.New()

	//loops (using config) of how often to check each google service

	cal.Check()

}
