package calendar

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

/*Check the list of calendars the service is authorized to access and convert
any events in the next 24 hours to JSON.*/
func (c *Calendar) Check() {
	calendarList, err := c.Service.CalendarList.List().Do()
	if err != nil {
		log.Fatalf("Could not retrieve calendars %v", err.Error())
	}

	//For each calendar google has a record of, load the events and ouput them
	for _, gcal := range calendarList.Items {

		today, tomorrow := limitDates()
		if ok := c.loadEvents(gcal.Id, today, tomorrow); !ok {
			break
		}
		c.Output()
	}
}

//Output a subset of details for a list of Google calendar events as a JSON string
func (c *Calendar) Output() {
	for _, ge := range c.Events {
		e := convert(ge)

		data, err := json.Marshal(e)
		if err != nil {
			log.Printf("Failed to marshal json: %v", err.Error())
			continue
		}
		log.Println(string(data))
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
