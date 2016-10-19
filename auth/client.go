package auth

import (
	calendar "google.golang.org/api/calendar/v3"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

//Service reads credentials from a local file and establishes a client with access
//to google's calendar API.
func Service() *calendar.Service {
	ctx := context.Background()

	ts, err := google.DefaultTokenSource(ctx, calendar.CalendarReadonlyScope)
	if err != nil {
		panic(err)
	}
	client := oauth2.NewClient(ctx, ts)
	calendar, err := calendar.New(client)
	if err != nil {
		panic(err)
	}
	return calendar
}
