package calendar

import (
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
