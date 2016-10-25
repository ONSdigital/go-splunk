package calendar

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCalendar(t *testing.T) {
	Convey("Given the current time", t, func() {
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
		c.Events = createEvents(eventNum, true)

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
