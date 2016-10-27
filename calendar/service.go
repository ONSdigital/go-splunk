package calendar

import (
	"log"
	"net/http"

	"github.com/ONSdigital/go-splunk/auth"
	gcal "google.golang.org/api/calendar/v3"
)

//Calendar holds a google calendar service connection
type Calendar struct {
	Service *gcal.Service
	Events  []*gcal.Event
}

//New reads credentials from a local file and establishes a client with access
//to google's calendar API.
func New() *Calendar {
	new := func(c *http.Client) (interface{}, error) {
		return gcal.New(c)
	}
	c := &Calendar{}
	c.Service = auth.Client(gcal.CalendarReadonlyScope, new).(*gcal.Service)

	return c
}

//Load the events in a 24 hour period for the given calendar from Google
func (c *Calendar) loadEvents(id string, earliest string, latest string) bool {
	events, err := c.queryEvents(id, earliest, latest)
	if err != nil {
		log.Printf("Unable to retrieve todays events. %v", err.Error())
		return false
	}

	if len(events.Items) == 0 {
		log.Println("There are no events today.")
		return false
	}

	c.Events = events.Items
	return true
}

func (c *Calendar) queryEvents(id string, today string, tomorrow string) (*gcal.Events, error) {
	return c.Service.Events.
		List(id).
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(today).
		TimeMax(tomorrow).
		OrderBy("startTime").
		Do()
}
