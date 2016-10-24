package calendar

import (
	"bytes"
	"io"
	//"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	gcal "google.golang.org/api/calendar/v3"
)

func TestCalendar(t *testing.T) {
	Convey("Given a time", t, func() {
		now := time.Now()

		Convey("When setMidnight is called", func() {
			s := setMidnight(now)

			Convey("The time should be set to midnight on today's date", func() {
				So(s, ShouldContainSubstring, "T00:00:00Z")
			})
		})
	})

	Convey("Given that limitDates is called", t, func() {
		start, end := limitDates()

		Convey("When comparing the dates", func() {
			s, err := time.Parse(time.RFC3339, start)
			if err != nil {
				panic(err)
			}
			e, err := time.Parse(time.RFC3339, end)
			if err != nil {
				panic(err)
			}
			diff := e.Sub(s)

			Convey("The dates should be 24 hours apart", func() {
				So(diff, ShouldEqual, time.Duration(24*time.Hour))
			})
		})
	})

	Convey("Given a Calendar with a list of 2 google events", t, func() {
		c := &Calendar{}
		eventNum := 2
		c.Events = createEvents(eventNum)
		Convey("When Output is called", func() {
			output := captureOutput(c.Output)
			Convey("The log should be using the structured format", func() {
				So(output, ShouldContainSubstring, "\"namespace\"")
			})
			Convey("The log should contain 2 lines", func() {
				count := strings.Count(output, "\"namespace\"")
				So(count, ShouldEqual, eventNum)
			})
		})
	})
}

func createEvents(num int) []*gcal.Event {
	list := []*gcal.Event{}
	date := "2016-10-24"
	for x := 0; x < num; x++ {

		creator := &gcal.EventCreator{
			DisplayName: "bob" + strconv.Itoa(x),
			Email:       "bob" + strconv.Itoa(x) + "@google.com",
		}
		start := &gcal.EventDateTime{
			Date:     date,
			DateTime: "",
		}

		end := &gcal.EventDateTime{
			Date:     date,
			DateTime: "",
		}

		item := &gcal.Event{
			ColorId:     "color" + strconv.Itoa(x),
			Creator:     creator,
			Description: "this is a test event" + strconv.Itoa(x),
			End:         end,
			Id:          "testevent" + strconv.Itoa(x),
			Kind:        "kind" + strconv.Itoa(x),
			Status:      "good" + strconv.Itoa(x),
			Start:       start,
			Summary:     "Event Test Appointment" + strconv.Itoa(x),
		}

		list = append(list, item)
	}
	return list
}

func captureOutput(f func()) string {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		os.Stdout = stdout
	}()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f()

	w.Close()
	out := <-outC
	return out
}
