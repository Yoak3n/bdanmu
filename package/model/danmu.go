package model

type DanMu struct {
	Sender    User   `json:"user"`
	Content   string `json:"content"`
	RoomId    int    `json:"room_id"`
	Type      int8   `json:"type"`
	MessageId string `json:"message_id"`
}
