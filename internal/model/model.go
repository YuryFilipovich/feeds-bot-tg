package model

import "time"

type Item struct { //Article feed element
	Title      string    `json:"title"`
	Categories []string  `json:"categories"` //Category for filtering
	Link       string    `json:"link"`
	Date       time.Time `json:"date"` //Source published date
	Summary    string    `json:"summary"`
	SourceName string    `json:"source_name"`
}

type Source struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	FeedURL   string    `json:"feed_url"`
	CreatedAt time.Time `json:"created_at"`
}

type Article struct {
	ID          int64     `json:"id"`
	SourceID    int64     `json:"source_id"`
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Summary     string    `json:"summary"`
	PublishedAt time.Time `json:"published_at"` //Source published
	PostedAt    time.Time `json:"posted_at"`    //Tg published
	CreatedAt   time.Time `json:"created_at"`
}
