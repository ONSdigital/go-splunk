package analytics

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/ONSdigital/go-ns/log"
	ga "google.golang.org/api/analytics/v3"
)

type query struct {
	Metrics    string `json:"metrics"`
	Start      string `json:"start-date,omitempty"`
	End        string `json:"end-date,omitempty"`
	Sort       string `json:"sort,omitempty"`
	Filters    string `json:"filters,omitempty"`
	Dimensions string `json:"dimensions,omitempty"`
	MaxResults string `json:"max-results,omitempty"`
}

type metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

//metric is formed of the details required for a single query to the GA API
type metric struct {
	Name      string   `json:"name"`
	Frequency string   `json:"frequency"`
	Query     query    `json:"query"`
	Meta      metadata `json:"meta"`
}

//metricDefinitions holds the list of all queries defined in the corresponding json file
type metricDefinitions struct {
	Metrics []*metric `json:"metrics"`
}

func (a *Analytics) processAnalytics(id string) {
	md := loadMetricsFile()
	for _, metric := range md.Metrics {
		q := metric.Query.build(id, a.Service.Data.Ga)

		results, err := q.Do()
		if err != nil {
			log.ErrorC("Invalid GA Request", err, map[string]interface{}{"request details": nil})
		}

		table := convert(results)

		//TODO:improve this for description and name being optional?
		log.Debug("GA - "+metric.Meta.Description, map[string]interface{}{"name": metric.Meta.Name, "results": table})
	}
}

func convert(results *ga.GaData) []string {
	var table []string

	c := len(results.ColumnHeaders)
	r := len(results.Rows)

	for i := 0; i < r; i++ {
		row := ""
		for j := 0; j < c; j++ {
			row += results.ColumnHeaders[j].Name + "=" + results.Rows[i][j] + " "
		}
		row = strings.TrimSpace(row)
		table = append(table, row)
	}

	return table
}

func (q *query) build(id string, service *ga.DataGaService) *ga.DataGaGetCall {
	call := service.Get("ga:"+id,
		q.Start,
		q.End,
		q.Metrics)

	if q.Dimensions != "" {
		call = call.Dimensions(q.Dimensions)
	}

	if q.Filters != "" {
		call = call.Filters(q.Filters)
	}

	if q.Sort != "" {
		call = call.Sort(q.Sort)
	}

	if q.MaxResults != "" {
		max, err := strconv.ParseInt(q.MaxResults, 0, 64)
		if err != nil {
			log.ErrorC("GA - cannot convert max results", err, map[string]interface{}{"max results": q.MaxResults})
		}
		call = call.MaxResults(max)
	}
	return call
}

func loadMetricsFile() *metricDefinitions {
	file, err := ioutil.ReadFile("googleAnalytics.json")
	if err != nil {
		log.ErrorC("GA - new request cannot read file", err, nil)
	}

	var d *metricDefinitions
	if err := json.Unmarshal(file, &d); err != nil {
		data := make(map[string]interface{})
		data["file contents"] = file
		log.ErrorC("Unmarshal gaFields.json", err, data)
	}

	return d
}
