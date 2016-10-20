package analytics

import (
	"net/http"

	"github.com/ONSdigital/go-splunk/auth"
	ga "google.golang.org/api/analytics/v3"
)

//Analytics holds a google analytics service connection
type Analytics ga.Service

//New reads credentials from a local file and establishes a client with access
//to google's analytics API.
func New() Analytics {
	fn := func(c *http.Client) (interface{}, error) {
		return ga.New(c)
	}
	c := auth.Client(ga.AnalyticsReadonlyScope, fn).(Analytics)

	return c
}
