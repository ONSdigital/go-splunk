package analytics

import (
	"net/http"
	"time"

	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-splunk/auth"
	"github.com/fatih/structs"
	ga "google.golang.org/api/analytics/v3"
)

//Analytics holds a google analytics service connection
type Analytics struct {
	Service *ga.Service
	Ticker  <-chan time.Time
}

//New reads credentials from a local file and establishes a client with access
//to google's analytics API.
func New(period *int) *Analytics {
	new := func(c *http.Client) (interface{}, error) {
		return ga.New(c)
	}

	a := &Analytics{}
	//may need to use AnalyticsManageUsersReadonlyScope temporarily to get ProfileIds the user has access to
	a.Service = auth.Client(ga.AnalyticsReadonlyScope, new).(*ga.Service)
	a.Ticker = time.NewTicker(time.Second * time.Duration(*period)).C

	return a
}

type Property struct {
	Name     string
	Profiles []string
}

//UPDATE THIS COMMENT ---
/*Check the list of calendars the service is authorized to access and convert
any events in the next 24 hours to JSON.*/
func (a *Analytics) Check() {
	summaries, err := a.Service.Management.AccountSummaries.List().Do()
	if err != nil {
		log.ErrorC("ga-management", err, nil)
		return
	}
	var properties []*Property

	for _, account := range summaries.Items { //for every account
		for _, propertySummary := range account.WebProperties { // for each property
			p := &Property{ // save the id of the property
				Name: propertySummary.Name,
			}
			for _, profileSummary := range propertySummary.Profiles { // and get an id for each view within that property
				p.Profiles = append(p.Profiles, profileSummary.Id)
			}
			properties = append(properties, p)
		}
	}

	//today, tomorrow := limitDates()

	for _, property := range properties {
		for _, profile := range property.Profiles {
			results := a.fetchData(profile)
			for _, d := range results {
				log.Debug("GA - "+properties[0].Name, structs.Map(d))
			}
		}
	}
}
