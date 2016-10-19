package cal

import "google.golang.org/api/calendar/v3"

//Event is a subset of the google calendar event for logging out human details
//rather than google metadata.
type Event struct {
	ColorID      string `json:"colorId,omitempty"`
	CreatorName  string `json:"creatorName,omitempty"`
	CreatorEmail string `json:"creatorEmail,omitempty"`
	Description  string `json:"description,omitempty"`
	End          string `json:"end,omitempty"`
	ID           string `json:"id,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Start        string `json:"start,omitempty"`
	Status       string `json:"status,omitempty"`
	Summary      string `json:"summary,omitempty"`
}

//Convert a google calendar event to a flattened smaller Event struct.
func Convert(item *calendar.Event) *Event {
	// If the DateTime is an empty string the Event is an all-day Event.
	// So only Date is available.
	start := item.Start.Date
	if item.Start.DateTime != "" {
		start = item.Start.DateTime
	}
	end := item.End.Date
	if item.End.DateTime != "" {
		end = item.End.DateTime
	}
	return &Event{
		ColorID:      item.ColorId,
		CreatorName:  item.Creator.DisplayName,
		CreatorEmail: item.Creator.Email,
		Description:  item.Description,
		End:          end,
		ID:           item.Id,
		Kind:         item.Kind,
		Start:        start,
		Status:       item.Status,
		Summary:      item.Summary,
	}
}
