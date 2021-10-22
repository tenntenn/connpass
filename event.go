package connpass

import "time"

// Series repressents an organization or event series.
type Series struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Type repressents type of event.
type Type string

const (
	TypeParticipation = "participation"
	TypeAdvertisement = "advertisement"
)

// Event repressents a single event of connpass.
type Event struct {
	ID               int       `json:"event_id"`
	Title            string    `json:"title"`
	Catch            string    `json:"catch"`
	Description      string    `json:"description"`
	URL              string    `json:"event_url"`
	Hashtag          string    `json:"hash_tag"`
	StartedAt        time.Time `json:"started_at"`
	EndedAt          time.Time `json:"ended_at"`
	Limit            int       `json:"limit"`
	Type             Type      `json:"event_type"`
	Series           Series    `json:"series"`
	Address          string    `json:"address"`
	Place            string    `json:"place"`
	Lat              string    `json:"lat"`
	Lon              string    `json:"lon"`
	OwnerID          int       `json:"owner_id"`
	OwnerNickname    string    `json:"owner_nickname"`
	OwnerDisplayName string    `json:"owner_display_name"`
	Accepted         int       `json:"accepted"`
	Waiting          int       `json:"waiting"`
	UpdatedAt        time.Time `json:"updated_at"`
}
