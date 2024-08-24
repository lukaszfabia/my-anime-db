package models

type Controllable interface {
	VoiceActor |
		Character |
		Anime |
		Post |
		User // add rest
}
