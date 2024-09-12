package models

// for orm

type Score string

type StatusAnime string

type WatchStatus string

type Pegi string

type AnimeType string

type CastRole string

type FriendRequestStatus string

const (
	Accepted FriendRequestStatus = "accepted"
	Rejected FriendRequestStatus = "rejected"
	Pending  FriendRequestStatus = "pending"
	Cancel   FriendRequestStatus = "cancel"
)

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
	Canceled    WatchStatus = "canceled"
	OnHold      WatchStatus = "on-hold"
	Completed   WatchStatus = "completed"
	PlanToWatch WatchStatus = "plan-to-watch"
	Watching    WatchStatus = "watching"
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

var AllAnimeStatuses = []StatusAnime{
	Finished,
	CurrentlyAiring,
	Unknown,
	Planned,
}

var AllScores = []Score{
	Bad,
	Boring,
	Average,
	Good,
	VeryGood,
	MasterPiece,
}

var AllWatchStatuses = []WatchStatus{
	Canceled,
	Completed,
	OnHold,
	PlanToWatch,
}

var AllPegis = []Pegi{
	PG13,
	R,
	G,
}

var AllCastRoles = []CastRole{
	Main,
	Episodic,
	Supporting,
}

var AllAnimeTypes = []AnimeType{
	TV,
	ONA,
	OVA,
	Movie,
}

var AllFriendRequestStatus = []FriendRequestStatus{
	Rejected,
	Pending,
	Accepted,
	Cancel,
}
