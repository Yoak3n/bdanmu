package model

type Room struct {
	User          `json:"user"`
	ShortId       int   `json:"short_id"`
	LongId        int64 `json:"long_id"`
	FollowerCount int64 `json:"follower_count"`
}

type Message struct {
	Type int `json:"type"`
	Data any `json:"data"`
}
