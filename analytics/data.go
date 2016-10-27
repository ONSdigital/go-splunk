package analytics

type Data struct {
	OS       string
	Browser  string
	UserType string
	Num      string
}

func (a *Analytics) fetchData(id string) []*Data {
	data, err := a.Service.Data.Ga.Get("ga:"+id, "yesterday", "today", "ga:sessions").
		Dimensions("ga:operatingSystem,ga:browser,ga:userType").
		Do()
	if err != nil {
		panic(err)
	}

	var d []*Data
	for _, r := range data.Rows {
		de := &Data{
			OS:       r[0],
			Browser:  r[1],
			UserType: r[2],
			Num:      r[3],
		}

		d = append(d, de)
	}

	return d
}
