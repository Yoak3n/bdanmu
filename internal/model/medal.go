package model

type Medal struct {
	Name     string `json:"name"`
	OwnerID  int64  `json:"owner_id"`
	Level    int    `json:"level,omitempty"`
	TargetID int64  `json:"target_id"`
	Color    int    `json:"color,omitempty"`
}
