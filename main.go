package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/onsdigital/gocal-to-json/auth"
	"github.com/onsdigital/gocal-to-json/event"
)

func main() {
	srv := auth.Service()

	cals, err := srv.CalendarList.List().Do()
	if err != nil {
		log.Fatalf("Could not retrieve calendars %v", err.Error())
	}

	today, tomorrow := limitDates()

	for _, c := range cals.Items {
		events, err := srv.Events.List(c.Id).ShowDeleted(false).
			SingleEvents(true).TimeMin(today).TimeMax(tomorrow).OrderBy("startTime").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve todays events. %v", err.Error())
		}

		if len(events.Items) == 0 {
			log.Println("There are no events today.")
			os.Exit(1)
		}

		for _, i := range events.Items {
			e := cal.Convert(i)

			data, err := json.Marshal(e)
			if err != nil {
				log.Fatalf("Failed to marshal json: %v", err.Error())
			}
			log.Println(string(data))
		}
	}

}

func limitDates() (today string, tomorrow string) {
	t := time.Now()
	today = setMidnight(t)
	tomorrow = setMidnight(t.Add(24 * time.Hour))
	return
}

func setMidnight(t time.Time) string {
	c := strings.Split(t.String(), " ")
	t, err := time.Parse(time.RFC3339, c[0]+"T00:00:00Z")
	if err != nil {
		panic(err)
	}
	return t.Format(time.RFC3339)
}
