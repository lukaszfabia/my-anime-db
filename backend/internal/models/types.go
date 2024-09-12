package models

type Controllable interface {
	VoiceActor |
		Character |
		Anime |
		Post |
		User |
		Genre |
		Studio
}
