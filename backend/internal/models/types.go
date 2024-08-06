package models

// for orm

type Score string

type StatusAnime string

type Status string

type Pegi string

type AnimeType string

type CastRole string

const (
	Bad         Score = "bad"
	Boring      Score = "boring"
	Average     Score = "average"
	Good        Score = "good"
	VeryGood    Score = "very-good"
	MasterPiece Score = "masterpiece"
)

const (
	Finished        StatusAnime = "finished"
	CurrentlyAiring StatusAnime = "currently-airing"
	Unknown         StatusAnime = "unknown"
	Planned         StatusAnime = "planned"
)

const (
	Canceled    Status = "canceled"
	OnHold      Status = "on-hold"
	Completed   Status = "completed"
	PlanToWatch Status = "plan-to-watch"
)

const (
	PG13 Pegi = "PG-13"
	R    Pegi = "R-17+"
	G    Pegi = "G-all"
)

const (
	TV    AnimeType = "tv"
	ONA   AnimeType = "ona"
	OVA   AnimeType = "ova"
	Movie AnimeType = "movie"
)

const (
	Main       CastRole = "main"
	Episodic   CastRole = "episodic"
	Supporting CastRole = "supporting"
)
