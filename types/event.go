package types

import "time"

type Event struct {
	ID             int64         `json:"id"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Date           time.Time     `json:"date"`
	Lat            float32       `json:"lat"`
	Lon            float32       `json:"lon"`
	IsPrivate      bool          `json:"is_private"`
	Interests      *AllInterests `json:"interests"`
	SameInterests  []*Interests  `json:"same_interests"`
	Owner          *User         `json:"owner"`
	Attendees      []*User       `json:"attendees"`
	IsUserAttended bool          `json:"is_user_attended"`
}

type EventsList struct {
	Attended  []*Event `json:"attended"`
	Private   []*Event `json:"private"`
	Suggested []*Event `json:"suggested"`
	OwnEvents []*Event `json:"own_events"`
}

func NewEventsList() *EventsList {
	return &EventsList{
		Attended:  make([]*Event, 0),
		Private:   make([]*Event, 0),
		Suggested: make([]*Event, 0),
		OwnEvents: make([]*Event, 0),
	}
}
