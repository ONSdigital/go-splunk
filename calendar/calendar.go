package calendar

import (
	"strings"
	"time"

	"github.com/ONSdigital/go-ns/log"
	"github.com/fatih/structs"
)

/*Check the list of calendars the service is authorized to access and convert
any events in the next 24 hours to JSON.*/
func (c *Calendar) Check() {
	calendarList, err := c.Service.CalendarList.List().Do()
	if err != nil {
		log.ErrorC("Could not retrieve calendars", err, nil)
	}

	today, tomorrow := limitDates()

	//For each calendar google has a record of, load the events and ouput them
	//TODO - IF CONFIG ITEM (SINGLE-CALENDAR) PROVIDED, ONLY RETURN RESULTS FROM THAT, NOT LOOP
	for _, gcal := range calendarList.Items {
		if ok := c.loadEvents(gcal.Id, today, tomorrow); !ok {
			break
		}
		c.Output()
	}
}

//Output a subset of details for a list of Google calendar events as a JSON string
func (c *Calendar) Output() {
	for _, googleEvent := range c.Events {
		customEvent := convert(googleEvent)
		log.Debug("Calendar event", structs.Map(customEvent))
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
		log.ErrorC("Failed to parse time", err, nil)
	}
	return t.Format(time.RFC3339)
}
