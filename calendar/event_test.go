package calendar

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	gcal "google.golang.org/api/calendar/v3"
)

func TestEvent(t *testing.T) {
	Convey("Given a google calendar event", t, func() {
		events := createEvents(1, false)
		ge := events[0]

		Convey("When it's an all day event", func() {
			date := "2016-10-24"
			ge.Start = &gcal.EventDateTime{
				Date:     date,
				DateTime: "",
			}

			ge.End = &gcal.EventDateTime{
				Date:     date,
				DateTime: "",
			}
			Convey("The converted Start and End fields should contain only a date", func() {
				event := convert(ge)
				So(event.Start, ShouldEqual, date)
			})
		})

		Convey("When it's an event with a start and end time", func() {
			date := "2016-10-24"
			ge.Start = &gcal.EventDateTime{
				DateTime: date + "T10:45:00-07:00",
			}

			ge.End = &gcal.EventDateTime{
				DateTime: date + "T11:00:00-07:00",
			}
			Convey("The converted Start and End fields should contain a timestamp", func() {
				event := convert(ge)
				So(event.Start, ShouldContainSubstring, "T10:45")
				So(event.End, ShouldContainSubstring, "T11:00")
			})
		})
	})
}

func createEvents(num int, withDates bool) []*gcal.Event {
	list := []*gcal.Event{}
	date := "2016-10-24"
	for x := 0; x < num; x++ {

		creator := &gcal.EventCreator{
			DisplayName: "bob" + strconv.Itoa(x),
			Email:       "bob" + strconv.Itoa(x) + "@google.com",
		}

		item := &gcal.Event{
			ColorId:     "color" + strconv.Itoa(x),
			Creator:     creator,
			Description: "this is a test event" + strconv.Itoa(x),
			Id:          "testevent" + strconv.Itoa(x),
			Kind:        "kind" + strconv.Itoa(x),
			Summary:     "Event Test Appointment" + strconv.Itoa(x),
		}

		if withDates {
			start := &gcal.EventDateTime{
				Date:     date,
				DateTime: "",
			}

			end := &gcal.EventDateTime{
				Date:     date,
				DateTime: "",
			}
			item.End = end
			item.Start = start
		}

		list = append(list, item)
	}
	return list
}
