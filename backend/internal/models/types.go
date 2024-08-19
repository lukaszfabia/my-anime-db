package models

// for orm

type Score string

type StatusAnime string

type Status string

type Pegi string

type AnimeType string

type CastRole string

type GenreOption string

type FriendRequestStatus string

const (
	Accepted FriendRequestStatus = "accepted"
	Rejected FriendRequestStatus = "rejected"
	Pending  FriendRequestStatus = "pending"
	Cancel   FriendRequestStatus = "cancel"
)

const (
	Action           GenreOption = "action"
	Cyberpunk        GenreOption = "cyberpunk"
	Drama            GenreOption = "drama"
	Ecchi            GenreOption = "ecchi"
	Experimental     GenreOption = "experimental"
	Fantasy          GenreOption = "fantasy"
	Harem            GenreOption = "harem"
	Hentai           GenreOption = "hentai"
	Historical       GenreOption = "historical"
	Horror           GenreOption = "horror"
	Comedy           GenreOption = "comedy"
	Crime            GenreOption = "crime"
	Magic            GenreOption = "magic"
	Mecha            GenreOption = "mecha"
	MaleHarem        GenreOption = "male-harem"
	Music            GenreOption = "music"
	Supernatural     GenreOption = "supernatural"
	Madness          GenreOption = "madness"
	SliceOfLife      GenreOption = "slice-of-life"
	Parody           GenreOption = "parody"
	Adventure        GenreOption = "adventure"
	Psychological    GenreOption = "psychological"
	Romance          GenreOption = "romance"
	RomanceSeparated GenreOption = "romance-separated"
	SciFi            GenreOption = "sci-fi"
	ShoujoAi         GenreOption = "shoujo-ai"
	ShounenAi        GenreOption = "shounen-ai"
	SpaceOpera       GenreOption = "space-opera"
	Sports           GenreOption = "sports"
	Steampunk        GenreOption = "steampunk"
	School           GenreOption = "school"
	MartialArts      GenreOption = "martial-arts"
	Mystery          GenreOption = "mystery"
	Thriller         GenreOption = "thriller"
	Military         GenreOption = "military"
	Yaoi             GenreOption = "yaoi"
	Yuri             GenreOption = "yuri"
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

var AllStatus = []Status{
	Canceled,
	OnHold,
	Completed,
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

var AllGenreOptions = []GenreOption{
	Action,
	Cyberpunk,
	Drama,
	Ecchi,
	Experimental,
	Fantasy,
	Harem,
	Hentai,
	Historical,
	Horror,
	Comedy,
	Crime,
	Magic,
	Mecha,
	MaleHarem,
	Music,
	Supernatural,
	Madness,
	SliceOfLife,
	Parody,
	Adventure,
	Psychological,
	Romance,
	RomanceSeparated,
	SciFi,
	ShoujoAi,
	ShounenAi,
	SpaceOpera,
	Sports,
	Steampunk,
	School,
	MartialArts,
	Mystery,
	Thriller,
	Military,
	Yaoi,
	Yuri,
}

var AllFriendRequestStatus = []FriendRequestStatus{
	Rejected,
	Pending,
	Accepted,
	Cancel,
}
