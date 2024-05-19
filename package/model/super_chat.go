package model

type SuperChat struct {
	User      User   `json:"user"`
	Content   string `json:"content"`
	RoomID    int    `json:"room_id"`
	MessageID string `json:"message_id"`
	Timestamp int    `json:"timestamp"`
	EndTime   int    `json:"end_time"`
	Price     int    `json:"price"`
}
